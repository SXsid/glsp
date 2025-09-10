import * as vscode from "vscode";
import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";

let client: LanguageClient;

export function activate(context: vscode.ExtensionContext) {
  // Path to your LSP server executable
  const serverCommand = "/home/shekhar/Personal/glsp/glsp";

  // Server options - how to start the LSP server
  const serverOptions: ServerOptions = {
    run: { command: serverCommand, transport: TransportKind.stdio },
    debug: { command: serverCommand, transport: TransportKind.stdio },
  };

  // Client options - configure the language client
  const clientOptions: LanguageClientOptions = {
    // Register the server for markdown files
    documentSelector: [{ scheme: "file", language: "markdown" }],
    synchronize: {
      // Notify the server about file changes to '.md' files contained in the workspace
      fileEvents: vscode.workspace.createFileSystemWatcher("**/*.md"),
    },
  };

  // Create the language client
  client = new LanguageClient(
    "glsp",
    "Custom GLSP Language Server",
    serverOptions,
    clientOptions,
  );

  // Start the client (this will also start the server)
  client
    .start()
    .then(() => {
      vscode.window.showInformationMessage(
        "GLSP Language Server started successfully!",
      );
    })
    .catch((error) => {
      vscode.window.showErrorMessage(
        `Failed to start GLSP Language Server: ${error.message}`,
      );
    });
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
