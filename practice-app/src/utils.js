export const shapes = ['Straight', 'Draw', 'Fade', 'Hook', 'Slice'];

export const obstructions = [
  'Under trees',
  'Over tree',
  'Around tree L',
  'Around tree R',
  'Bunker short',
  'Bunker L',
  'Bunker R',
  'OB L',
  'OB R'
];

export const clubs = {
  dr: 250,
  '3w': 230,
  '5w': 210,
  hybrid: 190,
  '5i': 170,
  '6i': 160,
  '7i': 150,
  '8i': 140,
  '9i': 130,
  pw: 120,
  gw: 100,
  sw: 80,
  lw: 70
};

export function generateRandomNumber(mode, shapes) {
  let randomYardage = null;
  let randomShape = null;
  let randomObstruction = null;

  if (mode === 'yardage-shape' || mode === 'course') {
    const shapeIndex = Math.floor(Math.random() * shapes.length);
    randomShape = shapes[shapeIndex];
  }

  if (mode === 'course') {
    const obstructionIndex = Math.floor(Math.random() * obstructions.length);
    randomObstruction = obstructions[obstructionIndex];
  }

  const interval = 5;
  const min = 30;
  const max = 270;
  const range = (max - min) / interval;
  const randomIndex = Math.floor(Math.random() * (range + 1));
  randomYardage = min + (randomIndex * interval);

  return { randomYardage, randomShape, randomObstruction };
}
