var shadow$provide = {};
(function(){
var b=this||self;
function c(a){var d=typeof a;if("object"==d)if(a){if(a instanceof Array)return"array";if(a instanceof Object)return d;var f=Object.prototype.toString.call(a);if("[object Window]"==f)return"object";if("[object Array]"==f||"number"==typeof a.length&&"undefined"!=typeof a.splice&&"undefined"!=typeof a.propertyIsEnumerable&&!a.propertyIsEnumerable("splice"))return"array";if("[object Function]"==f||"undefined"!=typeof a.call&&"undefined"!=typeof a.propertyIsEnumerable&&!a.propertyIsEnumerable("call"))return"function"}else return"null";else if("function"==
d&&"undefined"==typeof a.call)return"object";return d};function e(a){this.a=a}e.prototype.toString=function(){return this.a};function g(a){var d=arguments.length;if(1==d&&"array"==c(arguments[0]))return g.apply(null,arguments[0]);for(var f={},h=0;h<d;h++)f[arguments[h]]=!0;return f};g(new e("APPLET"),new e("BASE"),new e("EMBED"),new e("IFRAME"),new e("LINK"),new e("MATH"),new e("META"),new e("OBJECT"),new e("SCRIPT"),new e("STYLE"),new e("SVG"),new e("TEMPLATE"));"undefined"!==typeof Symbol&&c(Symbol);function k(){console.log("init");return console.log("start")}var l=["starter","browser","init"],m=b;l[0]in m||"undefined"==typeof m.execScript||m.execScript("var "+l[0]);for(var n;l.length&&(n=l.shift());)l.length||void 0===k?m=m[n]&&m[n]!==Object.prototype[n]?m[n]:m[n]={}:m[n]=k;
}).call(this);