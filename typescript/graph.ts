type Index = number;
type Weight = number;

export interface Graph {
    getAdjacent(node: Index): Index[];
    getWeight(u: Index, v: Index): Weight;
    size(): number;
}

// Adjacency matrix Graph
export class MatrixGraph implements Graph {
    private _edges: Weight[][];

    constructor(edges: Weight[][]) {
        this._edges = edges;
    }

    getAdjacent(node: Index): Index[] {
        return this._edges
            .map((_, i) => i)
            .filter((i) => this.getWeight(node, i) != Infinity);
    }

    getWeight(u: Index, v: Index): Weight {
        return this._edges[u][v];
    }

    size() {
        return this._edges.length;
    }
}

export type Node = {
    index: number;
    weight: number;
};

// Adjacency List Graph
export class ListGraph implements Graph {
    private _edges: Array<Node[]>;

    constructor(edges: Array<Node[]>) {
        this._edges = edges;
    }

    getAdjacent(node: number): number[] {
        return this._edges[node].map((node) => node.index);
    }

    getWeight(u: number, v: number): number {
        return this._edges[u].find((node) => node.index == v)?.weight || Infinity;
    }

    size(): number {
        return this._edges.length;
    }
}
