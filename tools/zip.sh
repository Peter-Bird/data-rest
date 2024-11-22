#!/bin/bash

# Define the name of the zip file
archive_name="../wf-dba.zip"

# List the files to include in the zip
files=(                   
    "../main.go"             
    "../pkg/config.go"       
    "../pkg/handles.go"      
    "../pkg/routes.go"   
    "../pkg/server.go"         
    "../pkg/services.go"     
)

# Zip the files
zip "$archive_name" "${files[@]}" > /dev/null

# Check if the zip command was successful
if [ $? -eq 0 ]; then
    echo "Files successfully zipped into $archive_name"
else
    echo "An error occurred while zipping the files."
    exit 1
fi
