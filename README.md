# golf

Goals

1. Make tracking stats easier
2. Enable visualizations for stats across rounds

# Flow

1. Play round
2. Generate course json
3. Generate Sheet from script

# Generate course objects

From root directory, run `python src/generate_course_json.py <course_name>`.
It is expected for there to be a `<course_name>.txt` to be present in the `src` directory.

The format of the `.txt` file is `<hole number>,<hole par>,<hole handicap>`
