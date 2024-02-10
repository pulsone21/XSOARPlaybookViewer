import Dagre from "@dagrejs/dagre";
import { useCallback, useEffect, useMemo, useState } from "react";
import ReactFlow, {
  ReactFlowProvider,
  useReactFlow,
  Background,
  Node,
  Edge,
  OnNodesChange,
  OnEdgesChange,
  applyNodeChanges,
  applyEdgeChanges,
} from "reactflow";

import { initialNodes, initialEdges } from "./nodes";
import "reactflow/dist/style.css";
import ConditionNode from "./Components/ConditionNode";
import AutomationNode from "./Components/AutomationNode";
import LabelNode from "./Components/LabelNode";
import StartNode from "./Components/StartNode";


const g = new Dagre.graphlib.Graph().setDefaultEdgeLabel(() => ({}));

const getLayoutedElements = (nodes: Node[], edges: Edge[], options: string) => {
  g.setGraph({ rankdir: options });

  edges.forEach((edge) => g.setEdge(edge.source, edge.target));
  // @ts-expect-error
  nodes.forEach((node) => g.setNode(node.id, node));

  Dagre.layout(g);

  return {
    nodes: nodes.map((node) => {
      const { x, y } = g.node(node.id);

      return { ...node, position: { x, y } };
    }),
    edges,
  };
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
    []
  );
  const { fitView } = useReactFlow();
  const [nodes, setNodes] = useState<Node[]>(initialNodes);
  const [edges, setEdges] = useState<Edge[]>(initialEdges);

  const onNodesChange: OnNodesChange = useCallback(
    (changes) => setNodes((nds) => applyNodeChanges(changes, nds)),
    [setNodes]
  );
  const onEdgesChange: OnEdgesChange = useCallback(
    (changes) => setEdges((eds) => applyEdgeChanges(changes, eds)),
    [setEdges]
  );

  useEffect(() => {
    const layouted = getLayoutedElements(nodes, edges, "TB");

    setNodes([...layouted.nodes]);
    setEdges([...layouted.edges]);

    window.requestAnimationFrame(() => {
      fitView();
    });
  }, [nodes, edges, setNodes, setEdges, fitView]);

  return (
    <ReactFlow
      nodes={nodes}
      edges={edges}
      onNodesChange={onNodesChange}
      onEdgesChange={onEdgesChange}
      fitView
      nodeTypes={nodeTypes}
      //@ts-expect-error
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
