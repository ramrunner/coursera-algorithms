import edu.princeton.cs.algs4.Digraph;
import edu.princeton.cs.algs4.Queue;
import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import java.util.Map;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Objects;

public class SAP {
	// constructor takes a digraph (not necessarily a DAG)
	private Digraph dg;
	private class nodepair {
		int v1;
		int v2;
		@Override 
		public boolean equals(Object o) {
			if (o == this) {
				return true;
			}
			if (!(o instanceof nodepair)) {
				return false;
			}
			nodepair no = (nodepair) o;
			return (no.v1 == v1 && no.v2 == v2);
		}
		@Override
		public int hashCode() {
			return Objects.hash(v1, v2);
		}
		private nodepair(int v, int w) {
			this.v1 = v;
			this.v2 = w;
		}
	}

	private class sapath {
		int length;
		int ancestor;
		private sapath(int length, int ancestor) {
			this.length = length;
			this.ancestor = ancestor;
		}
	}

	private class lockstepBfs {
		private boolean[] marked;
		private int[] distfromv1;
		private int[] distfromv2;
		private int v1,v2,ancestor;
		private boolean found, empty1, empty2, noancestor;
		private lockstepBfs(int v1, int v2) {
			this.v1 = v1;
			this.v2 = v2;
			this.ancestor = 0;
			this.found = false;
			this.empty1 = false;
			this.empty2 = false;
			this.noancestor = false;
		}

		private void bfs(Digraph g) {
			this.noancestor = true;
			Queue<Integer> q1 = new Queue<Integer>();
			Queue<Integer> q2 = new Queue<Integer>();
			marked = new boolean[g.V()];
			distfromv1 = new int[g.V()];
			distfromv2 = new int[g.V()];
			Arrays.fill(distfromv1,-1);
			Arrays.fill(distfromv2,-1);
			q1.enqueue(v1);
			q2.enqueue(v2);
			distfromv1[v1] = 0;
			distfromv2[v2] = 0;
			while (empty1!=true && empty2!=true) {
				//System.out.println("run");
				empty1 = dostep(q1,distfromv1,marked);
				if (!found) {
					empty2 = dostep(q2,distfromv2,marked);
				}
			}
		}
		//returns true when it's queue is empty
		private boolean dostep(Queue<Integer> q, int[] dist, boolean[] mark) {
			if (q.isEmpty()) {
				return true;
			}
			//System.out.println("stepping");
			int next = q.dequeue();
			if (mark[next] == true && dist[next] <= 0) {
				System.out.println(next+" is marked");
				found = true;
				this.noancestor = false;
				ancestor = next;
				//dist[next] = dist[next]+1;
				return false;
			}
			mark[next] = true;
			System.out.println("at node:"+next);
			for (int w : dg.adj(next)) {
				System.out.println("looking at node:"+w);
				if (mark[w] && dist[w] == -1) { //we found the ancestor
					found = true;
					this.noancestor = false;
					ancestor = w;
					dist[w] = dist[next]+1;
					System.out.println("ancestor dist of "+ w +":"+dist[w]);
					return false;
				}
				if (mark[w]) { //we have been here again
					//System.out.println("again");
					continue;
				}
				mark[w] = true;
				dist[w] = dist[next]+1;
				System.out.println("normal dist of "+ w +":"+dist[w]);
				q.enqueue(w);
			}
			return false;
		}
		private sapath getsapath() {
			sapath sp;
			if (this.noancestor) {
				sp = new sapath(-1,-1);
			} else {
				System.out.println("d1:"+distfromv1[ancestor]+" d2:"+distfromv2[ancestor]);
				sp = new sapath(distfromv1[ancestor]+distfromv2[ancestor],ancestor);
			}
			return sp;
		}
		
	}

	private Map<nodepair,sapath> spcache;

	public SAP(Digraph g) {
		dg = new Digraph(g);
		spcache = new HashMap<nodepair, sapath>();
	}

	// length of shortest ancestral path between v and w; -1 if no such path
	public int length(int v, int w) {
		if (v == w) {
			return 0;
		}
		sapath cacheget = spcache.get(new nodepair(v, w));
		if (cacheget != null) {
			return cacheget.length;
		}
		lockstepBfs lbfs = new lockstepBfs(v, w);
		lbfs.bfs(dg);
		sapath sp = lbfs.getsapath(); 
		spcache.put(new nodepair(v, w), sp);
		return sp.length;
	}

	// a common ancestor of v and w that participates in a shortest ancestral path; -1 if no such path
	public int ancestor(int v, int w) {
		if (v == w) {
			return v;
		}
		sapath cacheget = spcache.get(new nodepair(v, w));
		if (cacheget != null) {
			return cacheget.ancestor;
		}
		lockstepBfs lbfs = new lockstepBfs(v, w);
		lbfs.bfs(dg);
		sapath sp = lbfs.getsapath(); 
		spcache.put(new nodepair(v, w), sp);
		return sp.ancestor;

	}

	// length of shortest ancestral path between any vertex in v and any vertex in w; -1 if no such path
	public int length(Iterable<Integer> v, Iterable<Integer> w) {
		if (v == null || w == null) {
			throw new IllegalArgumentException("null v or w");
		}
		ArrayList<Integer> lens = new ArrayList<Integer>();
		for (Integer k: v) {
			for (Integer j: w) {
				if (k == null || j == null) {
					throw new IllegalArgumentException("null k or j");
				}
				lens.add(length(k, j));
			}
		}
		int smallest = -1;
		for (int k: lens) {
			if (k != -1) {
				if (smallest == -1) {
					smallest = k;//init it
				}
				if (k < smallest) {
					smallest = k;
				}
			}
		}
		return smallest;
	}

	// a common ancestor that participates in shortest ancestral path; -1 if no such path
	public int ancestor(Iterable<Integer> v, Iterable<Integer> w) {
		if (v == null || w == null) {
			throw new IllegalArgumentException("null v or w");
		}
		ArrayList<Integer> lens = new ArrayList<Integer>();
		ArrayList<Integer> ans = new ArrayList<Integer>();
		for (Integer k: v) {
			for (Integer j: w) {
				if (k == null || j == null) {
					throw new IllegalArgumentException("null k or j");
				}
				lens.add(length(k, j));
				ans.add(ancestor(k, j));
			}
		}
		int smallest = -1;
		for (int k: lens) {
			if (k != -1) {
				if (smallest == -1) {
					smallest = k;//init it
				}
				if (k < smallest) {
					smallest = k;
				}
			}
		}
		if (smallest == -1) {
			return -1;
		}
		return ans.get(smallest);

	}

	// do unit testing of this class
	public static void main(String[] args) {
		In in = new In(args[0]);
		Digraph G = new Digraph(in);
		SAP sap = new SAP(G);
		while (!StdIn.isEmpty()) {
			int v = StdIn.readInt();
			int w = StdIn.readInt();
			int length   = sap.length(v, w);
			int ancestor = sap.ancestor(v, w);
			StdOut.printf("length = %d, ancestor = %d\n", length, ancestor);
		}
	}
}
