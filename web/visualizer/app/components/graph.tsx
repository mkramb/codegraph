import React from "react"
import { GraphCanvas } from "reagraph"

interface DataItem {
  type: "label" | "link"
  value: string
  filepath: string
  position: {
    row: number
    column: number
  }
}

interface Node {
  id: string
  label: string
}

interface Edge {
  source: string
  target: string
  id: string
  label: string
}

export interface GraphComponentProps {
  data: DataItem[]
}

export const GraphComponent: React.FC<GraphComponentProps> = ({ data }) => {
  const parsedNodes: Node[] = []
  const parsedEdges: Edge[] = []

  const nodeIds = new Set<string>()

  data.forEach((item) => {
    if (item.type === "label") {
      if (!nodeIds.has(item.value)) {
        parsedNodes.push({
          id: item.value,
          label: item.value,
        })

        nodeIds.add(item.value)
      }
    } else if (item.type === "link") {
      const [source, target] = item.value.split(",")

      parsedEdges.push({
        source,
        target,
        id: `${source}-${target}`,
        label: `${source}-${target}`,
      })
    }
  })

  return <GraphCanvas nodes={parsedNodes} edges={parsedEdges} />
}
