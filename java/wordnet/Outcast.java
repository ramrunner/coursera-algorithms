import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.Collections;

public class Outcast {
	private final WordNet wn;
	private static class strdist {
		String noun;
		int    dst;
		private strdist(String n, int dst) {
			this.noun = n;
			this.dst = dst;
		}
	}
	private static class strdistCmp implements Comparator<strdist> {
		public int compare(strdist a, strdist b) {
			return a.dst - b.dst; 
		}
	}

	public Outcast(WordNet wn) {
		this.wn = wn;
	}

	public String outcast(String[] nouns) {
		ArrayList<strdist> dists = new ArrayList<strdist>();
		for(String i:nouns) {
			int ndst = 0, dst = 0;
			for (String j:nouns) {
				//if (i.equals(j)) {
				//	continue;
				//}
				dst = wn.distance(i, j);
				System.out.println(i+"->"+j+" dst:"+dst);
				ndst += dst;
			}
			dists.add(new strdist(i, ndst));
		}
		for (strdist i: dists) {
			System.out.println("noun:"+i.noun+" tdst:"+i.dst);
		}
		return Collections.max(dists, new strdistCmp()).noun;
	}

	public static void main(String[] args) {
		WordNet wn = new WordNet(args[0], args[1]);
		Outcast outcast = new Outcast(wn);
		for (int t = 2; t < args.length; t++) {
			In in = new In(args[t]);
			String[] nouns = in.readAllStrings();
			StdOut.println(args[t] + ": " + outcast.outcast(nouns));
		}
	}
}
