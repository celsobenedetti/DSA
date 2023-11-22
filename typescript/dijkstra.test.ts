import {expect, test} from 'vitest';
import {Dijkstra, getNearestUnvisited} from './dijkstra';
import {MatrixGraph} from './graph';

const _ = Infinity;

test('getNearestUnvisited', () => {
    const seen = new Set<number>();

    seen.add(0);
    seen.add(1);
    seen.add(2);

    const want = 3;
    const dist = [1, 2, 3, want, 5];

    const got = getNearestUnvisited(dist, seen);
    expect(got).toBe(want);
});

test('Dijkstra', () => {
    const edges = [
        [_, _, _, 2, _, _],
        [_, _, 1, _, _, _],
        [_, _, _, 2, _, _],
        [_, 1, _, _, 2, _],
        [_, _, _, _, _, 3],
        [3, _, _, _, _, _],
    ];

    const graph = new MatrixGraph(edges);
    const dist = Dijkstra(graph, 0);

    expect(dist[0]).toBe(0);
    expect(dist[1]).toBe(3);
    expect(dist[2]).toBe(4);
    expect(dist[3]).toBe(2);
    expect(dist[4]).toBe(4);
    expect(dist[5]).toBe(7);
});
