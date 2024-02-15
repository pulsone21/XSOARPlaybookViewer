const position = { x: 0, y: 0 };

export const initialNodes = [
  {
    id: "1",
    type: "input",
    data: { label: "input" },
    position,
  },
  {
    id: "2",
    data: { label: "node 2" },
    position,
  },
  {
    id: "2a",
    data: { label: "node 2a" },
    position,
  },
  {
    id: "2a-1",
    data: { label: "node 2a-1" },
    position,
    parentNode: "2a",
    extent: "parent",
  },
  {
    id: "2a-2",
    data: { label: "node 2a-2" },
    position,
    parentNode: "2a",
    extent: "parent",
  },
  {
    id: "2b",
    data: { label: "node 2b" },
    position,
  },
  {
    id: "2c",
    data: { label: "node 2c" },
    position,
  },
  {
    id: "2d",
    data: { label: "node 2d" },
    position,
  },
  {
    id: "3",
    data: { label: "node 3" },
    position,
  },
];

export const initialEdges = [
  { id: "e12", source: "1", target: "2", type: "smoothstep" },
  { id: "e13", source: "1", target: "3", type: "smoothstep" },
  { id: "e22a", source: "2", target: "2a", type: "smoothstep" },
  { id: "e22b", source: "2", target: "2b", type: "smoothstep" },
  { id: "e22c", source: "2", target: "2c", type: "smoothstep" },
  { id: "e2c2d", source: "2c", target: "2d", type: "smoothstep" },
  { id: "e58", source: "2a", target: "2a-1", type: "smoothstep" },
  { id: "e59", source: "2a", target: "2a-2", type: "smoothstep" },
];
