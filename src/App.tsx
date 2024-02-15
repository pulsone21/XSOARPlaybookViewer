import ELK, { LayoutOptions } from "elkjs/lib/elk.bundled";
import { useCallback, useLayoutEffect, useMemo, useState } from "react";
import ReactFlow, {
  Background,
  Edge,
  Node,
  OnEdgesChange,
  OnNodesChange,
  ReactFlowProvider,
  applyEdgeChanges,
  applyNodeChanges,
  useReactFlow,
} from "reactflow";

import "reactflow/dist/style.css";
import AutomationNode from "./Components/AutomationNode";
import ConditionNode from "./Components/ConditionNode";
import LabelNode from "./Components/LabelNode";
import StartNode from "./Components/StartNode";
import { initialEdges, initialNodes } from "./assets/test_nodes";

const elk = new ELK();

// Elk has a *huge* amount of options to configure. To see everything you can
// tweak check out:
//
// - https://www.eclipse.org/elk/reference/algorithms.html
// - https://www.eclipse.org/elk/reference/options.html
const elkOptions = {
  "elk.algorithm": "layered",
  "elk.layered.spacing.nodeNodeBetweenLayers": "100",
  "elk.spacing.nodeNode": "80",
};

const getLayoutedElements = (
  nodes: Node[],
  edges: Edge[],
  options: LayoutOptions,
) => {
  const isHorizontal = options?.["elk.direction"] === "RIGHT";
  const graph = {
    id: "root",
    layoutOptions: options,
    children: nodes.map((node) => ({
      ...node,
      // Adjust the target and source handle positions based on the layout
      // direction.
      targetPosition: isHorizontal ? "left" : "top",
      sourcePosition: isHorizontal ? "right" : "bottom",

      // Hardcode a width and height for elk to use when layouting.
      width: 150,
      height: 50,
    })),
    edges: edges,
  };

  return (
    elk
      // @ts-expect-error - Convertion issue from ReactFlow but works anyway
      .layout(graph)
      .then((layoutedGraph) => ({
        // @ts-expect-error - Don't know how to solve the possible undefined issue
        nodes: layoutedGraph.children.map((node) => ({
          ...node,
          // React Flow expects a position property on the node instead of `x`
          // and `y` fields.
          position: { x: node.x, y: node.y },
        })),
        edges: layoutedGraph.edges,
      }))
      .catch(console.error)
  );
};

// eslint-disable-next-line react-refresh/only-export-components
const LayoutFlow = () => {
  const nodeTypes = useMemo(
    () => ({
      condition: ConditionNode,
      automation: AutomationNode,
      label: LabelNode,
      start: StartNode,
    }),
    [],
  );
  const { fitView } = useReactFlow();
  const [nodes, setNodes] = useState<Node[]>(initialNodes);
  const [edges, setEdges] = useState<Edge[]>(initialEdges);

  const onLayout = useCallback(
    ({ direction, useInitialNodes = false }) => {
      const opts = { "elk.direction": direction, ...elkOptions };
      const ns: Node[] = useInitialNodes ? initialNodes : nodes;
      const es: Edge[] = useInitialNodes ? initialEdges : edges;

      getLayoutedElements(ns, es, opts).then(
        ({ nodes: layoutedNodes, edges: layoutedEdges }) => {
          setNodes(layoutedNodes);
          setEdges(layoutedEdges);

          window.requestAnimationFrame(() => fitView());
        },
      );
    },
    [nodes, edges, fitView],
  );

  // Calculate the initial layout on mount.
  useLayoutEffect(() => {
    onLayout({ direction: "DOWN", useInitialNodes: true });
  }, [onLayout]);

  return (
    <ReactFlow
      nodes={nodes}
      edges={edges}
      fitView
      nodeTypes={nodeTypes}
      //@ts-expect-error - IDK
      attributionPosition="hidden"
    >
      <Background />
    </ReactFlow>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export default function () {
  return (
    <ReactFlowProvider>
      <LayoutFlow />
    </ReactFlowProvider>
  );
}
