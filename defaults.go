package main

const DEFAULT_STYLE_SHEET = `html,body,div,span,applet,object,iframe,h1,h2,h3,h4,h5,h6,p,blockquote,pre,a,abbr,acronym,address,big,cite,code,del,dfn,em,img,ins,kbd,q,s,samp,small,strike,strong,sub,sup,tt,var,b,u,i,center,dl,dt,dd,ol,fieldset,form,label,legend,table,caption,tbody,tfoot,thead,tr,th,td,article,aside,canvas,details,embed,figure,figcaption,footer,header,hgroup,menu,nav,output,ruby,section,summary,time,mark,audio,video {margin:0;padding:0;border:0;font-size:100%;font:inherit;vertical-align:baseline;}
article,aside,details,figcaption,figure,footer,header,hgroup,menu,nav,section {display:block;}
body {line-height:1;}
blockquote,q {quotes:none;}
:root {--background:#eff1f5;--foreground:#4c4f69;--foreground-alt:#5c5f77;--primary-foreground:#1e66f5;--fs-xs:14px;--fs-sm:16px;--fs-md:18px;--fs-lg:20px;}
html,body {margin:0;padding:0;background:var(--background);color:var(--foreground);font-family:Arial,Helvetica,sans-serif;font-size:var(--fs-xs);line-height:1.5;}
#resume {max-width:1200px;margin:0 auto;}
a {color:var(--primary-foreground);text-decoration:none;transition:all 0.2s ease-in-out;font-weight:bold;}
h1 {text-align:center;font-size:var(--fs-lg);margin-bottom:0.8rem;}
h2 {font-size:var(--fs-md);margin-bottom:0.5rem;font-weight:bold;}
h3 {font-size:var(--fs-sm);color:var(--foreground-alt);margin-top:0.5rem;margin-bottom:0.5rem;font-weight:bold;}
.resume-section {padding:0.7rem;}
ul {list-style-type:disc;margin-top:0.5rem;}
li {margin-bottom:0.5rem;}
em {font-style:italic;}
strong {font-weight:bold;}`

const DEFAULT_HTML_TEMPLATE = `
  <!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width">
  <title>%s</title>
  <style>
    %s
  </style>
</head>
<body>
  <div id="resume">
    %s
  </div>
</body>
</html>
  `
