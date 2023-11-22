/*
 * Dijstra's shortest path algorithm
 *
 * Given a weighted - non negative graph of size N
 * */

import {Dijkstra} from './dijkstra';
import {MatrixGraph} from './graph';

const _ = Infinity;

const edges = [
    [_, _, _, 2, _, _],
    [_, _, 1, _, _, _],
    [_, _, _, 2, _, _],
    [_, 1, _, _, 2, _],
    [_, _, _, _, _, 3],
    [3, _, _, _, _, _],
];

const graph = new MatrixGraph(edges);

console.log({
    distances: Dijkstra(graph, 0),
});
