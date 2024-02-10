import React from "react";
import "../index.css";

export type ContainerProps = {
  children: React.ReactNode;
};
export default function ContainerNode({ children }: ContainerProps) {
  return (
    <div
      style={{
        border: "1px solid rgb(187, 187, 187)",
        padding: "8px",
        borderRadius: "4px",
        background: "white",
        boxShadow: "0px 0px 3px var(    --sdx-color-gray-tint-5)",
        pointerEvents: "none",
      }}
    >
      {children}
    </div>
  );
}

