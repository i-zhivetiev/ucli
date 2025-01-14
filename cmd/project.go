package cmd

import (
	"github.com/spf13/cobra"
)

type Project struct {
	PubKey               string   `json:"pub_key"`
	Name                 string   `json:"name"`
	IsSearchIndexAllowed bool     `json:"is_search_index_allowed"`
	IsSharedProject      bool     `json:"is_shared_project"`
	Features             Features `json:"features"`
}

type Features struct {
	GifToVideoConversion GifToVideoConversion `json:"gif_to_video_conversion"`
	MimeTypeFiltering    MimeTypeFiltering    `json:"mime_type_filtering"`
	TeamMembers          TeamMembers          `json:"team_members"`
	Uploads              Uploads              `json:"uploads"`
	VideoProcessing      IsEnabledField       `json:"video_processing"`
	MalwareProtection    IsEnabledField       `json:"malware_protection"`
	SvgValidation        IsEnabledField       `json:"svg_validation"`
}

type IsEnabledField struct {
	IsEnabled bool `json:"is_enabled"`
}

type GifToVideoConversion struct {
	IsEnabled bool `json:"is_enabled"`
}

type MimeTypeFiltering struct {
	MimeTypes  []string `json:"mime_types"`
	MimeGroups []string `json:"mime_groups"`
}

type TeamMembers struct {
	TeamSize int `json:"team_size"`
}

type Uploads struct {
	FilesizeLimit         int  `json:"filesize_limit"`
	Autostore             bool `json:"autostore"`
	ImageResolutionLimit  int  `json:"image_resolution_limit"`
	IsSignedUploadEnabled bool `json:"is_signed_upload_enabled"`
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "project management",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
