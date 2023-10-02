project_name="$1"

echo "Project name is $project_name"

mkdir "$project_name"

cp SAMPLE_README.md "$project_name"

echo "Created project $project_name"