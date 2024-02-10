import { Handle, Position } from "reactflow";
import { FaRegCirclePlay } from "react-icons/fa6";

export default function StartNode() {
  return (
    <>
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          background: "#7189a2",
          padding: "4px",
          borderRadius: "100px",
          border: "1px solid --sdx-color-gray",
        }}
      >
        <FaRegCirclePlay style={{ margin: "4px", color: "white" }} />
      </div>
      <Handle type="source" position={Position.Bottom} />
    </>
  );
}

