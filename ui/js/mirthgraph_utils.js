// Problems with resizing and jquery and chrome and this stuff is so dumb.
window.width = function() {
  return document.body.clientWidth;
};

window.height = function() {
  return document.body.clientHeight;
};

var createGraphFromData = function(data){
    // Create nodes
    console.info("Creating graph nodes");
    var nodes = {};
    for (var i = 0; i < data.nodes.length; i++) {
        var id = data.nodes[i].id;
        nodes[id] = new Node(id);
        nodes[id].report = {"Agent":[data.nodes[i].label]};
    }

    // Second link the nodes together
    console.info("Linking graph nodes");
    for (var i = 0; i < data.edges.length; i++){
    	try{
        	nodes[data.edges[i].source].addChild(nodes[data.edges[i].target]);
    	}catch(e){}
    	try{
	        nodes[data.edges[i].target].addParent(nodes[data.edges[i].source]);
    	}catch(e){}
    }


    // Create the graph and add the nodes
    var graph = new Graph();
    for (var id in nodes) {
        graph.addNode(nodes[id]);
    }
    
    console.log("Done creating graph from data");
    return graph;
}