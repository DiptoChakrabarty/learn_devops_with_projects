project_name="$1"

echo "Project name is $project_name"

mkdir "$project_name"

cp SAMPLE_README.md "$project_name/Readme.md"

echo "Created project $project_name"