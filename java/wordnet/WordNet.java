import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.Digraph;
import edu.princeton.cs.algs4.DirectedCycle;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.stream.Stream;
import java.util.stream.Collectors;
import java.lang.Integer;
import java.util.Collections;
import java.util.Comparator;
import java.util.NoSuchElementException;

public class WordNet {
   private ArrayList<syn> synsName,synsInd;
   private Digraph dg;
   private SAP sap;

   private class syn {
	   String name;
	   int ind;
	   public syn(String a, int b) {
		   this.name = a;
		   this.ind = b;
	   }
	   public String getName() {
		   return this.name;
	   }
	   public int getInd() {
		   return this.ind;
	   }
   }

   private static class synLexCmp implements Comparator<syn> {
	   public int compare(syn a, syn b) {
		   return a.getName().compareTo(b.getName()); //lexicographic
	   }
   }

   private static class synIndCmp implements Comparator<syn> {
   	   public int compare(syn a, syn b) {
	   	return a.ind - b.ind;
	   }
   }

   private static class dfsRoot {
   	   int root;
	   
	   public dfsRoot(Digraph g) {
	   	root = dfs(g, 12); //just a random start
	   }

	   public int dfs(Digraph g, int node) {
		Iterable<Integer> up = g.adj(node);
		if (up.spliterator().getExactSizeIfKnown() == 0) {
			return node; 
		}
		try {
			Integer first = up.iterator().next();
			return dfs(g, first); // always pick the first element to go up. sucks.
		} catch (NoSuchElementException e) {
			System.out.println("no such element:"+e);
			return node;
		}
	   }
	   
	   public int getRoot() {
	   	return root;
	   }
   }


   // constructor takes the name of the two input files
   public WordNet(String synsets, String hypernyms) {
   	   if (synsets == null || hypernyms == null) {
	   	throw new IllegalArgumentException("null arguments to constructor");
	   }
	   In synin = new In(synsets);
	   String str;
	   String[] parts, synparts;
	   Integer id = 0;
	   synsName = new ArrayList<syn>();
	   try {
		   while(true) {
			   str = synin.readLine();
			   if (str == null) {
			   	break;
			   }
			   //System.out.println("i read:"+str);
			   parts = str.split(",");
			   synparts = parts[1].split(" ");
			   id = Integer.parseInt(parts[0]);
			   for (String sp : synparts) {
			   	synsName.add(new syn(sp, id));
			   }
		   }
	   //create a digraph o id verts
	   dg = new Digraph(id+1);
	   synin.close();
	   synin = new In(hypernyms);
	          while(true) {
		  	str = synin.readLine();
			if (str == null) {
				break;
			}
			parts = str.split(",");
			id = Integer.parseInt(parts[0]);
			for (int j=1; j<parts.length;j++) {
				int oedge = Integer.parseInt(parts[j]);
				dg.addEdge(id, oedge);
			}
		  }
		  synin.close();
		  DirectedCycle dc = new DirectedCycle(dg);
		  if (dc.hasCycle()) {
		  	throw new IllegalArgumentException("cycle in directed graph");
		  }
		  Collections.sort(synsName, new synLexCmp());
		  this.sap = new SAP(dg);
		  this.synsInd = new ArrayList<>(synsName);
		  Collections.sort(synsInd, new synIndCmp());
	   } catch (NoSuchElementException e) {
		   System.out.println("EOF");
	   }
   }


   // returns all WordNet nouns
   public Iterable<String> nouns() {
	   return synsName.stream().map(s -> s.getName()).collect(Collectors.toList());
   }

   // is the word a WordNet noun?
   public boolean isNoun(String word) {
   	   int ind=0;
	   ind = Collections.binarySearch(synsName, new syn(word,-1), new synLexCmp());
	   if (ind >= 0) {
		//System.out.println("i found the syn:"+word);
	   	return true;
	   }
	   //System.out.println("syn:"+word+" not in the synlist");
	   return false;
   }

   private Iterable<Integer> getNoun(String w) {
   	int ind=0;
	boolean first=true;
	ArrayList<Integer> ret = new ArrayList<Integer>();
	ind = Collections.binarySearch(synsName, new syn(w,-1), new synLexCmp());
	for (int i=ind; w.equals(synsName.get(i).name);) {
		if (first && w.equals(synsName.get(i-1).name)) {
			i--;
			first = false;
		}
		//System.out.println("word:"+synsName.get(i).name+" ind:"+ synsName.get(i).ind+" after:"+ synsName.get(i+1).name+" before:"+ synsName.get(i-1).name);
		ret.add(synsName.get(i).ind);
		i++;
	}
	return ret;
   }

   private String getInd(int i) {
   	int ind=0;
	ind = Collections.binarySearch(synsInd, new syn("", i), new synIndCmp());
   	syn s = synsInd.get(ind);
	return s.name;
   }

   // distance between nounA and nounB (defined below)
   public int distance(String nounA, String nounB) {
	int d,dmin;
   	if (!isNoun(nounA) || !isNoun(nounB)) {
		throw new IllegalArgumentException("nouns not in wordnet");
	}
	dmin = d = -1;
	for(int i :getNoun(nounA)) {
		for(int j :getNoun(nounB)) {
			if (i == j) {
				continue;
			}
			d = sap.length(i,j);
			if (dmin == -1 && d >= 0) {
				dmin = d;
			}
			if (d < dmin) {
				dmin = d;
			}
		}
	}
	//System.out.println("calling length with i:"+i+" j:"+j);
	return dmin;

   }

   // a synset (second field of synsets.txt) that is the common ancestor of nounA and nounB
   // in a shortest ancestral path (defined below)
   public String sap(String nounA, String nounB) {
	int an,d,dmin,ind,indmin;
	ArrayList<Integer> ancs = new ArrayList<Integer>();
   	if (!isNoun(nounA) || !isNoun(nounB)) {
		throw new IllegalArgumentException("nouns not in wordnet");
	}
	ind = indmin = 0;
	d = dmin = -1;
	//System.out.println("chahch");
	for(int i :getNoun(nounA)) {
		//System.out.println("out with i:"+i);
		for(int j :getNoun(nounB)) {
			//System.out.println("in with j:"+j);
			if (i == j) {
				continue;
			}
			an = sap.ancestor(i,j);
			ancs.add(an);
			d = sap.length(i,j);
			if (dmin == -1 && d >= 0) {
				dmin = d;
				indmin = ind;
			}
			if (d < dmin) {
				dmin = d;
				indmin = ind;
			}
			ind++;
		}
	}
	System.out.println("ind is:"+indmin+"and arrlist:"+ancs);
	return getInd(ancs.get(indmin));
   }

   // do unit testing of this class
   public static void main(String[] args) {
   	   String[] syns = new String[]{"zero_hour", "zap", "flengifhidon"};
	   WordNet wn = new WordNet(args[0], args[1]);
	   System.out.println("finding the root...");
	   dfsRoot dr = new dfsRoot(wn.dg);
	   System.out.println("root is:"+dr.getRoot());
	   System.out.println("looking for syns...");
	   try {
		   for (String s : syns) {
			wn.isNoun(s);
		   }
	   } catch (IllegalArgumentException e) {
		System.out.println("correctly caught exception for nx noun");
	   }
	   System.out.println("doing worm and bird and the first bird should be "+wn.getInd(24306));
	   System.out.println("ancestor of worm and bird "+wn.sap("worm", "bird"));
	   System.out.println("distance of worm and bird "+wn.distance("worm", "bird"));
	   System.out.println("ancestor of marlin and mileage "+wn.sap("white_marlin", "mileage"));
	   System.out.println("distance of marlin and mileage "+wn.distance("white_marlin", "mileage"));
   }
}
