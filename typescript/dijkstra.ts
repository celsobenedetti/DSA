/*
 * Dijstra's shortest path algorithm
 *
 * Given a weighted - non negative graph of size N
 * */

import {Graph} from './graph';

export function Dijkstra(G: Graph, source: number): number[] {
    const N = G.size();
    const dist = Array(N).fill(Infinity);
    const seen = new Set<number>();

    dist[source] = 0;

    while (seen.size < N) {
        const current = getNearestUnvisited(dist, seen);

        seen.add(current);

        for (const neighbor of G.getAdjacent(current)) {
            const newDist = dist[current] + G.getWeight(current, neighbor);

            if (newDist < dist[neighbor]) {
                dist[neighbor] = newDist;
            }
        }
    }

    return dist;
}

export function getNearestUnvisited(dist: number[], seen: Set<number>): number {
    let min = Infinity;
    let result = 0;
    for (let i = 0; i < dist.length; i++) {
        if (!seen.has(i) && dist[i] < min) {
            min = dist[i];
            result = i;
        }
    }
    return result;
}
