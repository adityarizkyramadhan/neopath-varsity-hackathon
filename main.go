package main

import supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"

func main() {
	supClient := supabasestorageuploader.NewSupabaseClient(
		"PROJECT_URL",
		"PROJECT_API_KEYS",
		"STORAGE_NAME",
		"STORAGE_FOLDER",
	)

}
