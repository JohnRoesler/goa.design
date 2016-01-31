{
    "title": "goagen",
    "description": "The generator",
    "tags": [ "goagen", "dsl" ],
    "date": "2016-01-30",
    "categories": [
        "Generation"
    ]
}

 <p>
                                <code>goagen</code> is a tool that generates various artifacts from a goa design package.<br/> Install it with:
                                <code class="bash">go get github.com/goadesign/goa/goagen</code><br/>
                                goagen relies on <code>goimports</code>:<br/>
                                <code class="bash">go get golang.org/x/tools/cmd/goimports</code>
                        </p>
                        <p>
                                Each type of artifact is associated with a <code>goagen</code> command that exposes it own set of flags.
                                Internally these commands map to "generators" that contain the logic for generating the artifacts.
                                It works something like this:
                        </p>
                        <ol>
                                <li>
                                        goagen parses the command line to determine the type of output desired and invokes the appropriate generator.
                                </li>
                                <li>
                                        The generator writes the code of the tool that will produce the final output to a temporary Go workspace.
                                </li>
                                <li>
                                        The tool composed of the design language package and the output producing code is compiled in the temporary workspace.
                                </li>
                                <li>
                                        <code>goagen</code> then runs the tool which evaluates the design functions and traverses the resulting in-memory data
                                        structures to write the output.
                                </li>
                        </ol>
                        <p>
                                Each generator registers a command with the <code>goagen</code> tool, <code>goagen --help</code> lists all the available
                                commands. These are:
                        </p>
                        <ul>
                                <li>
                                        <code>app</code>: generates the service boilerplate code including controllers, contexts, media types and user types.
                                </li>
                                <li>
                                        <code>main</code>: generates a skeleton file for each resource controller as well as a default <code>main</code>.
                                </li>
                                <li>
                                        <code>client</code>: generates an API client Go package and tool.
                                </li>
                                <li>
                                        <code>swagger</code>: generates the API Swagger specification.
                                </li>
                                <li>
                                        <code>js</code>: generates a JavaScript API client.
                                </li>
                                <li>
                                        <code>schema</code>: generates the API Hyper-schema JSON.
                                </li>
                                <li>
                                        <code>gen</code>: invokes a third party generator.
                                </li>
                                <li>
                                        <code>bootstrap</code>: invokes the <code>app</code>, <code>main</code>, <code>client</code> and <code>swagger</code>
                                        generators.
                                </li>
                        </ul>
                        <p>
                                The command <code>goagen --help-long</code> lists all the supported commands and their flags.
                        </p>
                        <h2>Common flags</h2>
                        <p>
                                The following flags apply to all the <code>goagen</code> commands:
                        </p>
                        <ul>
                                <li>
                                        <code>--design|-d=DESIGN</code> defines the Go package path to the service design package.
                                </li>
                                <li>
                                        <code>--out|-o=OUT</code> specifies where to generate the files, defaults to the current directory.
                                </li>
                                        <li>
                                        <code>--debug</code> enables <code>goagen</code> debug. This causes <code>goagen</code> to print the content of the temporary
                                        files and to leave them around.
                                </li>
                                <li>
                                        <code>--help|--help-long|--help-man</code> prints contextual help.
                                </li>
                        </ul>
                        <h2><code>goagen app</code></h2>
                        <p>
                                The <code>app</code> command is arguably the most critical. It generates all the supporting code for the
                                goa service. This command supports an additional flag:
                        </p>
                        <ul>
                                <li>
                                        <code>--pkg=app</code> specifies the name of the generated Go package, defaults to <code>app</code>. That's also the
                                        name of the subdirectory that gets created to store the generated Go files.
                                </li>
                        </ul>
                        <p>
                                This command always deletes and re-creates any pre-existing directory with the same name. The idea
                                being that these files should never be edited.
                        </p>
                        <h2><code>goagen main</code></h2>
                        <p>
                                The <code>main</code> command helps bootstrap a new goa service by generating a default <code>main.go</code> as
                                well as a default (empty) implementation for each resource controller defined in the design package. By default
                                this command only generates the files if they don't exist yet in the output directory. This
                                command accepts two additional flags:
                        </p>
                        <ul>
                                <li>
                                        <code>--force</code> causes the files to be generated even if files with the same name already exist (in
                                        which case they get overwritten).
                                </li>
                                <li>
                                        <code>name=API</code> specifies the name of the service to be used in the generated call to <code>goa.New</code>.
                                </li>
                        </ul>
                        <h2><code>goagen client</code></h2>
                        <p>
                                The <code>client</code> command generates both an API client package and tool. The client package implements a <code>Client</code>
                                object that exposes one method for each resource action. The generated code of the CLI tool leverages the package to
                                make the API requests to the service.
                        </p>
                        <p>
                                The <code>Client</code> object can be configured to use request signers that get invoked prior to sending the
                                request. The signers modify the request to include auth headers for example. goa comes with signers that implement
                                <a href="https://godoc.org/github.com/goadesign/goa#BasicSigner">basic auth</a>,
                                <a href="https://godoc.org/github.com/goadesign/goa#JWTSigner">JWT auth</a> and a subset of
                                <a href="https://godoc.org/github.com/goadesign/goa#OAuth2Signer">OAuth2</a>. Custom signers my be used as well, they must
                                implement the <a href="https://godoc.org/github.com/goadesign/goa#Signer">Signer</a> interface.
                                This command accepts three additional flags:
                        </p>
                        <ul>
                                <li>
                                        <code>--version</code> specifies the CLI tool version.
                                </li>
                                <li>
                                        <code>--signer</code> specifies a signer object supported by the API. Signer objects generally perform auth.
                                </li>
                                <li>
                                        <code>--signerPkg</code> specifies the path to the package implementing the signer objects if not goa.
                                </li>
                        </ul>
                        <h2><code>goagen js</code></h2>
                        <p>
                                The <code>js</code> command generates a JavaScript API client suitable for both client-side and server-side
                                applications. The generated code defines an anonymous AMD module and relies on the <a href="https://github.com/mzabriskie/axios">axios</a>
                                promised-based JavaScript library for making the actual HTTP requests.
                        </p>
                        <p>
                                The generated module wraps the <code>axios</code> client and adds API specific functions, for example:
                        </p>
                        <pre><code class="javascript">// List all bottles in account optionally filtering by year
// path is the request path, the format is "/cellar/accounts/:accountID/bottles"
// years is used to build the request query string.
// config is an optional object to be merged into the config built by the function prior to making the request.
// The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
// This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
client.listBottle = function (path, years, config) {
  cfg = {
    timeout: timeout,
    url: urlPrefix + path,
    method: 'get',
    params: {
      years: years
    },
    responseType: 'json'
  };
  if (config) {
    cfg = utils.merge(cfg, config);
  }
  return client(cfg);
}</code></pre>
                        <p>
                                The generated client module can be loaded using <code>requirejs</code>:
                        </p>
                        <pre><code class="javascript">requirejs.config({
  paths: {
    axios: '/js/axios.min',
    client: '/js/client'
  }
});
requirejs(['client'], function (client) {
  client().listBottle ("/cellar/accounts/440/bottles", 317)
    .then(function (resp) {
      // All good, use resp
    })
    .catch(function (resp) {
      // Woops, request failed or returned 4xx or 5xx.
    });
});</code></pre>
                        <p>
                                The code above assumes that the generated files <code>client.js</code> and <code>axios.min.js</code> are both
                                served from <code>/js</code>. The <code>resp</code> value returned to the promise is an object with the following
                                fields:
                        </p>
                        <pre><code language="javascript">{
  // `data` is the response that was provided by the server
  data: {},

  // `status` is the HTTP status code from the server response
  status: 200,

  // `statusText` is the HTTP status message from the server response
  statusText: 'OK',

  // `headers` the headers that the server responded with
  headers: {},

  // `config` is the config that was provided to `axios` for the request
  config: {}
}</code></pre>
                        <p>
                                The generator also produces an example HTML and controller that can be mounted on a
                                goa service to quickly test the JavaScript. Simply import the <code>js</code> Go
                                package in your service main and mount the controller. The example HTML is served
                                under <code>/js</code> so that loading this path in a browser will trigger the generated
                                JavaScript.
                        </p>
                        <h2><code>goagen swagger</code></h2>
                        <p>
                                The <code>swagger</code> command generates a <a href="http://swagger.io">Swagger</a> specification of the API.
                                The command does not accept additional flags. It generates both the Swagger JSON as well as a controller that
                                can be mounted on the goa service to serve it under <code>/swagger.json</code>.
                        </p>
                        <h2><code>goagen schema</code></h2>
                        <p>
                                The <code>schema</code> command generates a <a href="https://blog.heroku.com/archives/2014/1/8/json_schema_for_heroku_platform_api">Heroku-like</a>
                                JSON hyper-schema representation of the API. It generates both the JSON as well as a controller that
                                can be mounted on the goa service to serve it under <code>/schema.json</code>. The command accepts an additional flag:
                        </p>
                        <ul>
                                <li>
                                        <code>--url|-u=URL</code> specifies the base URL used to build the JSON schema ID.
                                </li>
                        </ul>
                        <h2><code>goagen gen</code>: goagen Plugins</h2>
                        <p>
                                The <code>gen</code> command makes it possible to invoke <code>goagen</code> plugins.
                                This command accepts two flags:
                        </p>
                        <ul>
                                <li>
                                        <code>--pkg-path=PKG-PATH</code> specifies the Go package path to the plugin package.
                                </li>
                                <li>
                                        <code>--pkg-name=PKG-NAME</code> specifies the Go package name of the plugin package. It defaults to the
                                        name of the inner most directory in the Go package path.
                                </li>
                        </ul>
                        <p>
                                A plugin consists of a Go package which exposes a <code>Generate</code> function with the following
                                signature:
                        <p>
                        <pre><code class="go">func Generate(api *design.APIDefinition) ([]string, error)
</pre></code>
                        <p>
                                 where api is the API definition computed from the design package. On success <code>Generate</code> should return
                                 the path to the generated files. On error the error message gets displayed to the user (and
                                 goagen exist with status 1).
                        </p>
                        <p>
                                 The <code>Generate</code> method should take advantage of the <code>APIDefinition</code> <code>IterateXXX</code> methods to iterate
                                 through the API resources, media types and types to guarantee that the order doesn't change
                                 between two invokation of the function (thereby generating different output even if the design
                                 hasn't changed).
                        </p>
                        <p>
                                 They may also take advantage of Metadata. The goa design language allows defining Metadata on a
                                 number of definitions: API, Resource, Action, Response and Attribute (which means Type and
                                 MediaType as well since these definitions are attributes). A metadata field consists of a
                                 key/value pair where both are simple strings. The generator can use these key/value pairs to
                                 produce different results, see example below. Metadata has no effect on the buit-in generators.
                                 The Output directory is available through the codegen.OutputDir global variable.
                        </p>
                        <p>
                                 Package genresnames is an example of a goagen plugin. It creates a file "names.txt" containing
                                 the names of the API resources sorted in alphabetical order. If a resource has a
                                 metadata pair with the key "genresnames/name" then the plugin uses the metadata value instead.
                        </p>
                        <p>
                                 Invoke the plugin with:
                        </p>
                        <pre><code>
                                 goagen gen -d <Go package path to design package> --pkg-path=<Go package path to genresnames>
                        </code></pre>
                        <p>
                                 Source code:
                        </p>
                        <pre><code class="go">
package genresnames

import (
        "io/ioutil"
        "os"
        "path/filepath"
        "strings"

        "gopkg.in/alecthomas/kingpin.v2"

        "github.com/goadesign/goa/design"
        "github.com/goadesign/goa/goagen/codegen"
)

// Generate is the function called by goagen to generate the names file.
func Generate(api *design.APIDefinition) ([]string, error) {
        // Make sure to parse the common flags so that codegen.OutputDir gets properly
        // initialized.
        app := kingpin.New("Resource names", "Resource name generator")
        codegen.RegisterFlags(app)
        if _, err := app.Parse(os.Args[1:]); err != nil {
                panic(err)
        }

        // Now iterate through the resources to gather their names
        names := make([]string, len(api.Resources))
        i := 0
        api.IterateResources(func(res *design.ResourceDefinition) error {
                if n, ok := res.Metadata["genresnames/name"]; ok {
                        names[i] = n
                } else {
                        names[i] = res.Name
                }
                i++
                return nil
        })
        content := strings.Join(names, "\n")

        // Write the output file and return its name
        outputFile := filepath.Join(codegen.OutputDir, "names.txt")
        ioutil.WriteFile(outputFile, []byte(content), 0755)
        return []string{outputFile}, nil
}
</code></pre>

