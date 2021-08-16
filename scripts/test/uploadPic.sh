curl localhost:10000/query \
-H auth='7c5a08e2ce9f0038786e414ef2062fada8feb76f' \
-F operations=                                 '{ "query": "mutation($req: [UploadFile!]!) { multipleUpload(req: $req) { id, name, content } }", "variables": { "req": [ { "id": 1, "file": null }, { "id": 2, "file": null } ] } }' \
-F map='{ "0": ["variables.req.0.file"] }' \
-F 0=./j.jpeg

# curl 'http://localhost:10000/' --data-binary '{"query":"mutation($image: Upload!) {\n  createPost(\n    input: { title: \"mamad title\", body: \"mamads body\", image: $image }\n  )\n}\n","variables":{"image":"2"}}' --compressed