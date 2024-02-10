import { Handle, NodeProps, Position } from "reactflow";
import ContainerNode from "./BaseNode";
import { GoRepoForked } from "react-icons/go";

export type ConditionNodeData = {
  label: string;
  conditions?: string[];
};

export default function ConditionNode({ data }: NodeProps<ConditionNodeData>) {
  if (data.conditions === undefined) {
    data.conditions = ["default"];
  }

  return (
    <ContainerNode>
      <Handle type="target" position={Position.Top} />
      <div style={{ display: "flex", flexDirection: "row" }}>
        <p className="text-align-center h5 margin-right-1 margin-bottom-0">
          {data.label}
        </p>
        <GoRepoForked style={{ color: "#7189a2", alignSelf: "center" }} />
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "center",
        }}
      >
        {data.conditions.map((el, i) => {
          const conLen = data.conditions ? data.conditions.length : 2;
          const computedW = (100 / (conLen + 1)) * (i + 1);
          const w = computedW.toString().concat("%");
          return (
            <div key={el}>
              <p className="text-smaller margin-h-1 margin-bottom-0">{el}</p>
              <Handle
                type="source"
                id={el}
                position={Position.Bottom}
                style={{ left: w }}
              />
            </div>
          );
        })}
      </div>
    </ContainerNode>
  );
}

