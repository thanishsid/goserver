package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/db"
)

const TempVideoSignature = "temp_video_*"

type Video struct {
	DB db.DB
}

var _ domain.VideoService = (*Video)(nil)

// Reads the video data from reader, saves it to disk and returns video object.
func (v *Video) Save(ctx context.Context, input domain.VideoUploadInput) (*domain.Video, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// Check database to find if video with the same filename exists.
	fileExists, err := v.DB.CheckVideoFileExists(ctx, input.FileName)
	if err != nil {
		return nil, err
	}

	// Return ErrFileAlreadyExists if file exists.
	if fileExists {
		return nil, ErrFileAlreadyExists
	}

	// create a random uuid to use as video filename and the id for database entry.
	videoID := uuid.New()

	// Create a new temporary file in the video directory to write the file data.
	tempFile, err := os.CreateTemp(config.C.VideoDirectory, TempVideoSignature)
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	// Write the input file data to the temp File.
	written, err := io.Copy(tempFile, input.File)
	if err != nil {
		return nil, err
	}

	if written != input.Size {
		return nil, ErrIncompleteFile
	}

	// New id for the video thumbnail
	thumbnailID := uuid.New()

	// if input custom thumbnail is valid then override the current thumbnailID with it.
	if input.CustomThumbnailID.Valid {
		thumbnailID = input.CustomThumbnailID.UUID
	}

	// If a custom thumbnail is not provided run ffmpegThumbnailer to generate a video thumbnail
	// using the temp video file and output it to the image directory.
	if !input.CustomThumbnailID.Valid {
		videoDuration, err := getVideoDuration(tempFile.Name())
		if err != nil {
			return nil, err
		}

		// Get the thumbnail posiiton at 10% of the video duration.
		thumbnailPosition := int(videoDuration * 0.1)

		thumbnailFileName := fmt.Sprintf("vid_thumbnail_%s.jpg", videoID.String())

		// Generate the tumbnail.
		thumbnailCmd := exec.Command(
			"ffmpeg",
			"-ss", fmt.Sprint(thumbnailPosition),
			"-i", tempFile.Name(),
			"-vframes", "1",
			getImagePath(thumbnailFileName))
		if err := thumbnailCmd.Run(); err != nil {
			return nil, err
		}

		// thumbnail create time.
		now := time.Now()

		if err := v.DB.InsertOrUpdateImage(ctx, db.InsertOrUpdateImageParams{
			ID:        thumbnailID,
			FileName:  thumbnailFileName,
			CreatedAt: now,
			UpdatedAt: now,
		}); err != nil {
			return nil, err
		}
	}

	// video create time.
	now := time.Now()

	newVideo := domain.Video{
		ID:          videoID,
		FileName:    input.FileName,
		ThumbnailID: thumbnailID,
		Link:        config.C.VideoLink,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Insert video details to the database.
	if err := v.DB.InsertOrUpdateVideo(ctx, db.InsertOrUpdateVideoParams{
		ID:          newVideo.ID,
		FileName:    newVideo.FileName,
		ThumbnailID: newVideo.ThumbnailID,
		CreatedAt:   newVideo.CreatedAt,
		UpdatedAt:   newVideo.UpdatedAt,
	}); err != nil {
		return nil, err
	}

	// Rename the temp file name to regular video name.
	if err := os.Rename(tempFile.Name(), getVideoPath(newVideo.FileName)); err != nil {
		return nil, err
	}

	return &newVideo, nil
}

// Delete a video.
func (v *Video) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := v.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Get target video from database.
	video, err := tx.GetVideoById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	// Get thhumbnail of the target video.
	thumbnail, err := tx.GetImageById(ctx, video.ThumbnailID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	// Delete video info from database.
	if err := tx.DeleteVideo(ctx, video.ID); err != nil {
		return err
	}

	// Delete thumbnail image info from database.
	if err := tx.DeleteImage(ctx, thumbnail.ID); err != nil {
		return err
	}

	// Delete video file from disk.
	if os.Remove(getVideoPath(video.FileName)); err != nil {
		return err
	}

	// Delete thumbnail file from disk.
	if os.Remove(getImagePath(thumbnail.FileName)); err != nil {
		return err
	}

	// Commit transaction.
	return tx.Commit(ctx)
}

// Get a video.
func (v *Video) One(ctx context.Context, id uuid.UUID) (*domain.Video, error) {
	dbVideo, err := v.DB.GetVideoById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &domain.Video{
		ID:          dbVideo.ID,
		FileName:    dbVideo.FileName,
		ThumbnailID: dbVideo.ThumbnailID,
		Link:        config.C.VideoLink,
		CreatedAt:   dbVideo.CreatedAt,
		UpdatedAt:   dbVideo.UpdatedAt,
	}, nil
}

// Get Videos by ids.
func (v *Video) AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*domain.Video, error) {
	videoRows, err := v.DB.GetAllVideosInIDS(ctx, ids)
	if err != nil {
		return nil, err
	}

	videos := make([]*domain.Video, len(videoRows))

	for i, row := range videoRows {
		videos[i] = &domain.Video{
			ID:          row.ID,
			FileName:    row.FileName,
			ThumbnailID: row.ThumbnailID,
			Link:        config.C.VideoLink,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		}
	}

	return videos, nil
}
