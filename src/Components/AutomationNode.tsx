import { Handle, NodeProps, Position } from "reactflow";
import "../index.css";
import ContainerNode from "./BaseNode";
import { BsGearWideConnected } from "react-icons/bs";

export type AutomationNodeData = {
  label: string;
  args?: string[];
};

export default function AutomationNode({
  data,
}: NodeProps<AutomationNodeData>) {
  return (
    <ContainerNode>
      <Handle type="target" position={Position.Top} />
      <div style={{ display: "flex", flexDirection: "row" }}>
        <div className="h5  margin-right-2 margin-bottom-0 text-align-center">
          {data.label}
        </div>
        <BsGearWideConnected
          style={{ alignSelf: "center", color: "#7189a2" }}
        />
      </div>
      <div>
        {data.args &&
          data.args.map((el) => {
            return (
              <div key={el}>
                <p className="text-smaller">{el}</p>
              </div>
            );
          })}
      </div>
      <Handle type="source" position={Position.Bottom} />
    </ContainerNode>
  );
}

