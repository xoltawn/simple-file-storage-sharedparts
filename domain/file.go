package domain

// File is the domain objects for stored files
type File struct {
	//OriginalUrl indicates the url from which the file was downloaded(used when file is downloaded from a link)
	OriginalURL string `json:"original_url"`
	//LocalName is the name given on storing
	LocalName string `json:"local_name"`
	//FileExtension ...
	FileExtension string `json:"file_extension"`
	//FileSize ...
	FileSize int64 `json:"file_size"`
	//CreatedAt ...
	CreatedAt string `json:"created_at"`
}
