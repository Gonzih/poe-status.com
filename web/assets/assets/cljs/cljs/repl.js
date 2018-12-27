// Compiled by ClojureScript 1.10.339 {}
goog.provide('cljs.repl');
goog.require('cljs.core');
goog.require('cljs.spec.alpha');
cljs.repl.print_doc = (function cljs$repl$print_doc(p__18234){
var map__18235 = p__18234;
var map__18235__$1 = ((((!((map__18235 == null)))?(((((map__18235.cljs$lang$protocol_mask$partition0$ & (64))) || ((cljs.core.PROTOCOL_SENTINEL === map__18235.cljs$core$ISeq$))))?true:false):false))?cljs.core.apply.call(null,cljs.core.hash_map,map__18235):map__18235);
var m = map__18235__$1;
var n = cljs.core.get.call(null,map__18235__$1,new cljs.core.Keyword(null,"ns","ns",441598760));
var nm = cljs.core.get.call(null,map__18235__$1,new cljs.core.Keyword(null,"name","name",1843675177));
cljs.core.println.call(null,"-------------------------");

cljs.core.println.call(null,[cljs.core.str.cljs$core$IFn$_invoke$arity$1((function (){var temp__5720__auto__ = new cljs.core.Keyword(null,"ns","ns",441598760).cljs$core$IFn$_invoke$arity$1(m);
if(cljs.core.truth_(temp__5720__auto__)){
var ns = temp__5720__auto__;
return [cljs.core.str.cljs$core$IFn$_invoke$arity$1(ns),"/"].join('');
} else {
return null;
}
})()),cljs.core.str.cljs$core$IFn$_invoke$arity$1(new cljs.core.Keyword(null,"name","name",1843675177).cljs$core$IFn$_invoke$arity$1(m))].join(''));

if(cljs.core.truth_(new cljs.core.Keyword(null,"protocol","protocol",652470118).cljs$core$IFn$_invoke$arity$1(m))){
cljs.core.println.call(null,"Protocol");
} else {
}

if(cljs.core.truth_(new cljs.core.Keyword(null,"forms","forms",2045992350).cljs$core$IFn$_invoke$arity$1(m))){
var seq__18237_18259 = cljs.core.seq.call(null,new cljs.core.Keyword(null,"forms","forms",2045992350).cljs$core$IFn$_invoke$arity$1(m));
var chunk__18238_18260 = null;
var count__18239_18261 = (0);
var i__18240_18262 = (0);
while(true){
if((i__18240_18262 < count__18239_18261)){
var f_18263 = cljs.core._nth.call(null,chunk__18238_18260,i__18240_18262);
cljs.core.println.call(null,"  ",f_18263);


var G__18264 = seq__18237_18259;
var G__18265 = chunk__18238_18260;
var G__18266 = count__18239_18261;
var G__18267 = (i__18240_18262 + (1));
seq__18237_18259 = G__18264;
chunk__18238_18260 = G__18265;
count__18239_18261 = G__18266;
i__18240_18262 = G__18267;
continue;
} else {
var temp__5720__auto___18268 = cljs.core.seq.call(null,seq__18237_18259);
if(temp__5720__auto___18268){
var seq__18237_18269__$1 = temp__5720__auto___18268;
if(cljs.core.chunked_seq_QMARK_.call(null,seq__18237_18269__$1)){
var c__4351__auto___18270 = cljs.core.chunk_first.call(null,seq__18237_18269__$1);
var G__18271 = cljs.core.chunk_rest.call(null,seq__18237_18269__$1);
var G__18272 = c__4351__auto___18270;
var G__18273 = cljs.core.count.call(null,c__4351__auto___18270);
var G__18274 = (0);
seq__18237_18259 = G__18271;
chunk__18238_18260 = G__18272;
count__18239_18261 = G__18273;
i__18240_18262 = G__18274;
continue;
} else {
var f_18275 = cljs.core.first.call(null,seq__18237_18269__$1);
cljs.core.println.call(null,"  ",f_18275);


var G__18276 = cljs.core.next.call(null,seq__18237_18269__$1);
var G__18277 = null;
var G__18278 = (0);
var G__18279 = (0);
seq__18237_18259 = G__18276;
chunk__18238_18260 = G__18277;
count__18239_18261 = G__18278;
i__18240_18262 = G__18279;
continue;
}
} else {
}
}
break;
}
} else {
if(cljs.core.truth_(new cljs.core.Keyword(null,"arglists","arglists",1661989754).cljs$core$IFn$_invoke$arity$1(m))){
var arglists_18280 = new cljs.core.Keyword(null,"arglists","arglists",1661989754).cljs$core$IFn$_invoke$arity$1(m);
if(cljs.core.truth_((function (){var or__3949__auto__ = new cljs.core.Keyword(null,"macro","macro",-867863404).cljs$core$IFn$_invoke$arity$1(m);
if(cljs.core.truth_(or__3949__auto__)){
return or__3949__auto__;
} else {
return new cljs.core.Keyword(null,"repl-special-function","repl-special-function",1262603725).cljs$core$IFn$_invoke$arity$1(m);
}
})())){
cljs.core.prn.call(null,arglists_18280);
} else {
cljs.core.prn.call(null,((cljs.core._EQ_.call(null,new cljs.core.Symbol(null,"quote","quote",1377916282,null),cljs.core.first.call(null,arglists_18280)))?cljs.core.second.call(null,arglists_18280):arglists_18280));
}
} else {
}
}

if(cljs.core.truth_(new cljs.core.Keyword(null,"special-form","special-form",-1326536374).cljs$core$IFn$_invoke$arity$1(m))){
cljs.core.println.call(null,"Special Form");

cljs.core.println.call(null," ",new cljs.core.Keyword(null,"doc","doc",1913296891).cljs$core$IFn$_invoke$arity$1(m));

if(cljs.core.contains_QMARK_.call(null,m,new cljs.core.Keyword(null,"url","url",276297046))){
if(cljs.core.truth_(new cljs.core.Keyword(null,"url","url",276297046).cljs$core$IFn$_invoke$arity$1(m))){
return cljs.core.println.call(null,["\n  Please see http://clojure.org/",cljs.core.str.cljs$core$IFn$_invoke$arity$1(new cljs.core.Keyword(null,"url","url",276297046).cljs$core$IFn$_invoke$arity$1(m))].join(''));
} else {
return null;
}
} else {
return cljs.core.println.call(null,["\n  Please see http://clojure.org/special_forms#",cljs.core.str.cljs$core$IFn$_invoke$arity$1(new cljs.core.Keyword(null,"name","name",1843675177).cljs$core$IFn$_invoke$arity$1(m))].join(''));
}
} else {
if(cljs.core.truth_(new cljs.core.Keyword(null,"macro","macro",-867863404).cljs$core$IFn$_invoke$arity$1(m))){
cljs.core.println.call(null,"Macro");
} else {
}

if(cljs.core.truth_(new cljs.core.Keyword(null,"repl-special-function","repl-special-function",1262603725).cljs$core$IFn$_invoke$arity$1(m))){
cljs.core.println.call(null,"REPL Special Function");
} else {
}

cljs.core.println.call(null," ",new cljs.core.Keyword(null,"doc","doc",1913296891).cljs$core$IFn$_invoke$arity$1(m));

if(cljs.core.truth_(new cljs.core.Keyword(null,"protocol","protocol",652470118).cljs$core$IFn$_invoke$arity$1(m))){
var seq__18241_18281 = cljs.core.seq.call(null,new cljs.core.Keyword(null,"methods","methods",453930866).cljs$core$IFn$_invoke$arity$1(m));
var chunk__18242_18282 = null;
var count__18243_18283 = (0);
var i__18244_18284 = (0);
while(true){
if((i__18244_18284 < count__18243_18283)){
var vec__18245_18285 = cljs.core._nth.call(null,chunk__18242_18282,i__18244_18284);
var name_18286 = cljs.core.nth.call(null,vec__18245_18285,(0),null);
var map__18248_18287 = cljs.core.nth.call(null,vec__18245_18285,(1),null);
var map__18248_18288__$1 = ((((!((map__18248_18287 == null)))?(((((map__18248_18287.cljs$lang$protocol_mask$partition0$ & (64))) || ((cljs.core.PROTOCOL_SENTINEL === map__18248_18287.cljs$core$ISeq$))))?true:false):false))?cljs.core.apply.call(null,cljs.core.hash_map,map__18248_18287):map__18248_18287);
var doc_18289 = cljs.core.get.call(null,map__18248_18288__$1,new cljs.core.Keyword(null,"doc","doc",1913296891));
var arglists_18290 = cljs.core.get.call(null,map__18248_18288__$1,new cljs.core.Keyword(null,"arglists","arglists",1661989754));
cljs.core.println.call(null);

cljs.core.println.call(null," ",name_18286);

cljs.core.println.call(null," ",arglists_18290);

if(cljs.core.truth_(doc_18289)){
cljs.core.println.call(null," ",doc_18289);
} else {
}


var G__18291 = seq__18241_18281;
var G__18292 = chunk__18242_18282;
var G__18293 = count__18243_18283;
var G__18294 = (i__18244_18284 + (1));
seq__18241_18281 = G__18291;
chunk__18242_18282 = G__18292;
count__18243_18283 = G__18293;
i__18244_18284 = G__18294;
continue;
} else {
var temp__5720__auto___18295 = cljs.core.seq.call(null,seq__18241_18281);
if(temp__5720__auto___18295){
var seq__18241_18296__$1 = temp__5720__auto___18295;
if(cljs.core.chunked_seq_QMARK_.call(null,seq__18241_18296__$1)){
var c__4351__auto___18297 = cljs.core.chunk_first.call(null,seq__18241_18296__$1);
var G__18298 = cljs.core.chunk_rest.call(null,seq__18241_18296__$1);
var G__18299 = c__4351__auto___18297;
var G__18300 = cljs.core.count.call(null,c__4351__auto___18297);
var G__18301 = (0);
seq__18241_18281 = G__18298;
chunk__18242_18282 = G__18299;
count__18243_18283 = G__18300;
i__18244_18284 = G__18301;
continue;
} else {
var vec__18250_18302 = cljs.core.first.call(null,seq__18241_18296__$1);
var name_18303 = cljs.core.nth.call(null,vec__18250_18302,(0),null);
var map__18253_18304 = cljs.core.nth.call(null,vec__18250_18302,(1),null);
var map__18253_18305__$1 = ((((!((map__18253_18304 == null)))?(((((map__18253_18304.cljs$lang$protocol_mask$partition0$ & (64))) || ((cljs.core.PROTOCOL_SENTINEL === map__18253_18304.cljs$core$ISeq$))))?true:false):false))?cljs.core.apply.call(null,cljs.core.hash_map,map__18253_18304):map__18253_18304);
var doc_18306 = cljs.core.get.call(null,map__18253_18305__$1,new cljs.core.Keyword(null,"doc","doc",1913296891));
var arglists_18307 = cljs.core.get.call(null,map__18253_18305__$1,new cljs.core.Keyword(null,"arglists","arglists",1661989754));
cljs.core.println.call(null);

cljs.core.println.call(null," ",name_18303);

cljs.core.println.call(null," ",arglists_18307);

if(cljs.core.truth_(doc_18306)){
cljs.core.println.call(null," ",doc_18306);
} else {
}


var G__18308 = cljs.core.next.call(null,seq__18241_18296__$1);
var G__18309 = null;
var G__18310 = (0);
var G__18311 = (0);
seq__18241_18281 = G__18308;
chunk__18242_18282 = G__18309;
count__18243_18283 = G__18310;
i__18244_18284 = G__18311;
continue;
}
} else {
}
}
break;
}
} else {
}

if(cljs.core.truth_(n)){
var temp__5720__auto__ = cljs.spec.alpha.get_spec.call(null,cljs.core.symbol.call(null,[cljs.core.str.cljs$core$IFn$_invoke$arity$1(cljs.core.ns_name.call(null,n))].join(''),cljs.core.name.call(null,nm)));
if(cljs.core.truth_(temp__5720__auto__)){
var fnspec = temp__5720__auto__;
cljs.core.print.call(null,"Spec");

var seq__18255 = cljs.core.seq.call(null,new cljs.core.PersistentVector(null, 3, 5, cljs.core.PersistentVector.EMPTY_NODE, [new cljs.core.Keyword(null,"args","args",1315556576),new cljs.core.Keyword(null,"ret","ret",-468222814),new cljs.core.Keyword(null,"fn","fn",-1175266204)], null));
var chunk__18256 = null;
var count__18257 = (0);
var i__18258 = (0);
while(true){
if((i__18258 < count__18257)){
var role = cljs.core._nth.call(null,chunk__18256,i__18258);
var temp__5720__auto___18312__$1 = cljs.core.get.call(null,fnspec,role);
if(cljs.core.truth_(temp__5720__auto___18312__$1)){
var spec_18313 = temp__5720__auto___18312__$1;
cljs.core.print.call(null,["\n ",cljs.core.str.cljs$core$IFn$_invoke$arity$1(cljs.core.name.call(null,role)),":"].join(''),cljs.spec.alpha.describe.call(null,spec_18313));
} else {
}


var G__18314 = seq__18255;
var G__18315 = chunk__18256;
var G__18316 = count__18257;
var G__18317 = (i__18258 + (1));
seq__18255 = G__18314;
chunk__18256 = G__18315;
count__18257 = G__18316;
i__18258 = G__18317;
continue;
} else {
var temp__5720__auto____$1 = cljs.core.seq.call(null,seq__18255);
if(temp__5720__auto____$1){
var seq__18255__$1 = temp__5720__auto____$1;
if(cljs.core.chunked_seq_QMARK_.call(null,seq__18255__$1)){
var c__4351__auto__ = cljs.core.chunk_first.call(null,seq__18255__$1);
var G__18318 = cljs.core.chunk_rest.call(null,seq__18255__$1);
var G__18319 = c__4351__auto__;
var G__18320 = cljs.core.count.call(null,c__4351__auto__);
var G__18321 = (0);
seq__18255 = G__18318;
chunk__18256 = G__18319;
count__18257 = G__18320;
i__18258 = G__18321;
continue;
} else {
var role = cljs.core.first.call(null,seq__18255__$1);
var temp__5720__auto___18322__$2 = cljs.core.get.call(null,fnspec,role);
if(cljs.core.truth_(temp__5720__auto___18322__$2)){
var spec_18323 = temp__5720__auto___18322__$2;
cljs.core.print.call(null,["\n ",cljs.core.str.cljs$core$IFn$_invoke$arity$1(cljs.core.name.call(null,role)),":"].join(''),cljs.spec.alpha.describe.call(null,spec_18323));
} else {
}


var G__18324 = cljs.core.next.call(null,seq__18255__$1);
var G__18325 = null;
var G__18326 = (0);
var G__18327 = (0);
seq__18255 = G__18324;
chunk__18256 = G__18325;
count__18257 = G__18326;
i__18258 = G__18327;
continue;
}
} else {
return null;
}
}
break;
}
} else {
return null;
}
} else {
return null;
}
}
});

//# sourceMappingURL=repl.js.map
