type Index = number;
type Weight = number;

export interface Graph {
    getAdjacent(node: Index): Index[];
    getWeight(node1: Index, node2: Index): Weight;
    size(): number;
}

// Adjacency matrix Graph
export class MatrixGraph implements Graph {
    private _edges: Weight[][];

    constructor(edges: Weight[][]) {
        this._edges = edges;
    }

    getAdjacent(node: Index): Index[] {
        return this._edges.map((_, i) => i).filter(i => this.getWeight(node, i) != Infinity);
    }

    getWeight(u: Index, v: Index): Weight {
        return this._edges[u][v];
    }

    size() {
        return this._edges.length;
    }
}
