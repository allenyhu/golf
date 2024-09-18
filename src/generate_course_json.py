import json
import sys

def generate_course(course_name):
  holes = []

  with open(f"{course_name}.txt", 'r') as file:
    for line in file:
      hole_data = line.strip().split(',', 3)

      hole_number = hole_data[0]
      hole_par = hole_data[1]
      hole_handicap = hole_data[2]

      holes.append({'number': hole_number, 'handicap': hole_handicap, 'par': hole_par})
  
  with open(f"courses/{course_name}.json", 'w') as json_file:
    course = {'name': course_name, 'holes': holes}
    json.dump(course, json_file, indent=4)


generate_course(sys.argv[1])
