import { Handle, NodeProps, Position } from "reactflow";

export type LabelNodeData = {
  label: string;
};

export default function LabelNode({ data }: NodeProps<LabelNodeData>) {
  return (
    <>
      <Handle type="target" position={Position.Top} />
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          background: "#7189a2",
          padding: "4px",
          borderRadius: "12px",
          border: "1px solid --sdx-color-gray",
        }}
      >
        <p className="text-align-center text-small" style={{ color: "white" }}>
          {data.label}
        </p>
      </div>
      <Handle type="source" position={Position.Bottom} />
    </>
  );
}

