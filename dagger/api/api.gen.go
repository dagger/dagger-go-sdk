// Code generated by dagger. DO NOT EDIT.

package api

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"go.dagger.io/dagger/sdk/go/dagger/querybuilder"
)

// New returns a new API query object
func New(c graphql.Client) *Query {
	return &Query{
		q: querybuilder.Query(),
		c: c,
	}
}

// A global cache volume identifier
type CacheID string

// The address (also known as "ref") of a container published as an OCI image.
//
// Examples:
//   - "alpine"
//   - "index.docker.io/alpine"
//   - "index.docker.io/alpine:latest"
//   - "index.docker.io/alpine:latest@sha256deadbeefdeadbeefdeadbeef"
type ContainerAddress string

// A unique container identifier. Null designates an empty container (scratch).
type ContainerID string

// A content-addressed directory identifier
type DirectoryID string

type FileID string

// A unique identifier for a secret
type SecretID string

// Additional options for executing a command
type ExecOpts struct {
	// Optionally redirect the command's standard error to a file in the container.
	// Null means discard output.
	RedirectStderr string `json:"redirectStderr"`

	// Optionally redirect the command's standard output to a file in the container.
	// Null means discard output.
	RedirectStdout string `json:"redirectStdout"`

	// Optionally write to the command's standard input
	//
	// - Null means don't touch stdin (no redirection)
	// - Empty string means inject zero bytes to stdin, then send EOF
	Stdin string `json:"stdin"`
}

// A directory whose contents persist across runs
type CacheVolume struct {
	q *querybuilder.Selection
	c graphql.Client
}

func (r *CacheVolume) ID(ctx context.Context) (CacheID, error) {
	q := r.q.Select("id")

	var response CacheID
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// An OCI-compatible container, also known as a docker container
type Container struct {
	q *querybuilder.Selection
	c graphql.Client
}

// Default arguments for future commands
func (r *Container) DefaultArgs(ctx context.Context) ([]string, error) {
	q := r.q.Select("defaultArgs")

	var response []string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// Retrieve a directory at the given path. Mounts are included.
func (r *Container) Directory(path string) *Directory {
	q := r.q.Select("directory")
	q = q.Arg("path", path)

	return &Directory{
		q: q,
		c: r.c,
	}
}

// Entrypoint to be prepended to the arguments of all commands
func (r *Container) Entrypoint(ctx context.Context) ([]string, error) {
	q := r.q.Select("entrypoint")

	var response []string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// ContainerExecOpts contains options for Container.Exec
type ContainerExecOpts struct {
	Args []string

	Opts ExecOpts
}

// This container after executing the specified command inside it
func (r *Container) Exec(opts ...ContainerExecOpts) *Container {
	q := r.q.Select("exec")
	// `args` optional argument
	for i := len(opts) - 1; i >= 0; i-- {
		if !querybuilder.IsZeroValue(opts[i].Args) {
			q = q.Arg("args", opts[i].Args)
			break
		}
	}
	// `opts` optional argument
	for i := len(opts) - 1; i >= 0; i-- {
		if !querybuilder.IsZeroValue(opts[i].Opts) {
			q = q.Arg("opts", opts[i].Opts)
			break
		}
	}

	return &Container{
		q: q,
		c: r.c,
	}
}

// Exit code of the last executed command. Zero means success.
// Null if no command has been executed.
func (r *Container) ExitCode(ctx context.Context) (int, error) {
	q := r.q.Select("exitCode")

	var response int
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// Retrieve a file at the given path. Mounts are included.
func (r *Container) File(path string) *File {
	q := r.q.Select("file")
	q = q.Arg("path", path)

	return &File{
		q: q,
		c: r.c,
	}
}

// Initialize this container from the base image published at the given address
func (r *Container) From(address ContainerAddress) *Container {
	q := r.q.Select("from")
	q = q.Arg("address", address)

	return &Container{
		q: q,
		c: r.c,
	}
}

// A unique identifier for this container
func (r *Container) ID(ctx context.Context) (ContainerID, error) {
	q := r.q.Select("id")

	var response ContainerID
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// List of paths where a directory is mounted
func (r *Container) Mounts(ctx context.Context) ([]string, error) {
	q := r.q.Select("mounts")

	var response []string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// Publish this container as a new image
func (r *Container) Publish(ctx context.Context, address ContainerAddress) (ContainerAddress, error) {
	q := r.q.Select("publish")
	q = q.Arg("address", address)

	var response ContainerAddress
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// This container's root filesystem. Mounts are not included.
func (r *Container) Rootfs() *Directory {
	q := r.q.Select("rootfs")

	return &Directory{
		q: q,
		c: r.c,
	}
}

// The error stream of the last executed command.
// Null if no command has been executed.
func (r *Container) Stderr() *File {
	q := r.q.Select("stderr")

	return &File{
		q: q,
		c: r.c,
	}
}

// The output stream of the last executed command.
// Null if no command has been executed.
func (r *Container) Stdout() *File {
	q := r.q.Select("stdout")

	return &File{
		q: q,
		c: r.c,
	}
}

// The user to be set for all commands
func (r *Container) User(ctx context.Context) (string, error) {
	q := r.q.Select("user")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// The value of the specified environment variable
func (r *Container) Variable(ctx context.Context, name string) (string, error) {
	q := r.q.Select("variable")
	q = q.Arg("name", name)

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// A list of environment variables passed to commands
func (r *Container) Variables(ctx context.Context) ([]string, error) {
	q := r.q.Select("variables")

	var response []string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// ContainerWithDefaultArgsOpts contains options for Container.WithDefaultArgs
type ContainerWithDefaultArgsOpts struct {
	Args []string
}

// Configures default arguments for future commands
func (r *Container) WithDefaultArgs(opts ...ContainerWithDefaultArgsOpts) *Container {
	q := r.q.Select("withDefaultArgs")
	// `args` optional argument
	for i := len(opts) - 1; i >= 0; i-- {
		if !querybuilder.IsZeroValue(opts[i].Args) {
			q = q.Arg("args", opts[i].Args)
			break
		}
	}

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container but with a different command entrypoint
func (r *Container) WithEntrypoint(args []string) *Container {
	q := r.q.Select("withEntrypoint")
	q = q.Arg("args", args)

	return &Container{
		q: q,
		c: r.c,
	}
}

// ContainerWithMountedCacheOpts contains options for Container.WithMountedCache
type ContainerWithMountedCacheOpts struct {
	Source DirectoryID
}

// This container plus a cache volume mounted at the given path
func (r *Container) WithMountedCache(cache CacheID, path string, opts ...ContainerWithMountedCacheOpts) *Container {
	q := r.q.Select("withMountedCache")
	q = q.Arg("cache", cache)
	q = q.Arg("path", path)
	// `source` optional argument
	for i := len(opts) - 1; i >= 0; i-- {
		if !querybuilder.IsZeroValue(opts[i].Source) {
			q = q.Arg("source", opts[i].Source)
			break
		}
	}

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container plus a directory mounted at the given path
func (r *Container) WithMountedDirectory(path string, source DirectoryID) *Container {
	q := r.q.Select("withMountedDirectory")
	q = q.Arg("path", path)
	q = q.Arg("source", source)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container plus a file mounted at the given path
func (r *Container) WithMountedFile(path string, source FileID) *Container {
	q := r.q.Select("withMountedFile")
	q = q.Arg("path", path)
	q = q.Arg("source", source)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container plus a secret mounted into a file at the given path
func (r *Container) WithMountedSecret(path string, source SecretID) *Container {
	q := r.q.Select("withMountedSecret")
	q = q.Arg("path", path)
	q = q.Arg("source", source)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container plus a temporary directory mounted at the given path
func (r *Container) WithMountedTemp(path string) *Container {
	q := r.q.Select("withMountedTemp")
	q = q.Arg("path", path)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container plus an env variable containing the given secret
func (r *Container) WithSecretVariable(name string, secret SecretID) *Container {
	q := r.q.Select("withSecretVariable")
	q = q.Arg("name", name)
	q = q.Arg("secret", secret)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container but with a different command user
func (r *Container) WithUser(name string) *Container {
	q := r.q.Select("withUser")
	q = q.Arg("name", name)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container plus the given environment variable
func (r *Container) WithVariable(name string, value string) *Container {
	q := r.q.Select("withVariable")
	q = q.Arg("name", name)
	q = q.Arg("value", value)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container but with a different working directory
func (r *Container) WithWorkdir(path string) *Container {
	q := r.q.Select("withWorkdir")
	q = q.Arg("path", path)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container after unmounting everything at the given path.
func (r *Container) WithoutMount(path string) *Container {
	q := r.q.Select("withoutMount")
	q = q.Arg("path", path)

	return &Container{
		q: q,
		c: r.c,
	}
}

// This container minus the given environment variable
func (r *Container) WithoutVariable(name string) *Container {
	q := r.q.Select("withoutVariable")
	q = q.Arg("name", name)

	return &Container{
		q: q,
		c: r.c,
	}
}

// The working directory for all commands
func (r *Container) Workdir(ctx context.Context) (string, error) {
	q := r.q.Select("workdir")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// A directory
type Directory struct {
	q *querybuilder.Selection
	c graphql.Client
}

// DirectoryContentsOpts contains options for Directory.Contents
type DirectoryContentsOpts struct {
	Path string
}

// Return a list of files and directories at the given path
func (r *Directory) Contents(ctx context.Context, opts ...DirectoryContentsOpts) ([]string, error) {
	q := r.q.Select("contents")
	// `path` optional argument
	for i := len(opts) - 1; i >= 0; i-- {
		if !querybuilder.IsZeroValue(opts[i].Path) {
			q = q.Arg("path", opts[i].Path)
			break
		}
	}

	var response []string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// The difference between this directory and an another directory
func (r *Directory) Diff(other DirectoryID) *Directory {
	q := r.q.Select("diff")
	q = q.Arg("other", other)

	return &Directory{
		q: q,
		c: r.c,
	}
}

// Retrieve a directory at the given path
func (r *Directory) Directory(path string) *Directory {
	q := r.q.Select("directory")
	q = q.Arg("path", path)

	return &Directory{
		q: q,
		c: r.c,
	}
}

// Retrieve a file at the given path
func (r *Directory) File(path string) *File {
	q := r.q.Select("file")
	q = q.Arg("path", path)

	return &File{
		q: q,
		c: r.c,
	}
}

// The content-addressed identifier of the directory
func (r *Directory) ID(ctx context.Context) (DirectoryID, error) {
	q := r.q.Select("id")

	var response DirectoryID
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// load a project's metadata
func (r *Directory) LoadProject(configPath string) *Project {
	q := r.q.Select("loadProject")
	q = q.Arg("configPath", configPath)

	return &Project{
		q: q,
		c: r.c,
	}
}

// This directory plus the contents of the given file copied to the given path
func (r *Directory) WithCopiedFile(path string, source FileID) *Directory {
	q := r.q.Select("withCopiedFile")
	q = q.Arg("path", path)
	q = q.Arg("source", source)

	return &Directory{
		q: q,
		c: r.c,
	}
}

// This directory plus a directory written at the given path
func (r *Directory) WithDirectory(directory DirectoryID, path string) *Directory {
	q := r.q.Select("withDirectory")
	q = q.Arg("directory", directory)
	q = q.Arg("path", path)

	return &Directory{
		q: q,
		c: r.c,
	}
}

// DirectoryWithNewFileOpts contains options for Directory.WithNewFile
type DirectoryWithNewFileOpts struct {
	Contents string
}

// This directory plus a new file written at the given path
func (r *Directory) WithNewFile(path string, opts ...DirectoryWithNewFileOpts) *Directory {
	q := r.q.Select("withNewFile")
	// `contents` optional argument
	for i := len(opts) - 1; i >= 0; i-- {
		if !querybuilder.IsZeroValue(opts[i].Contents) {
			q = q.Arg("contents", opts[i].Contents)
			break
		}
	}
	q = q.Arg("path", path)

	return &Directory{
		q: q,
		c: r.c,
	}
}

// This directory with the directory at the given path removed
func (r *Directory) WithoutDirectory(path string) *Directory {
	q := r.q.Select("withoutDirectory")
	q = q.Arg("path", path)

	return &Directory{
		q: q,
		c: r.c,
	}
}

// This directory with the file at the given path removed
func (r *Directory) WithoutFile(path string) *Directory {
	q := r.q.Select("withoutFile")
	q = q.Arg("path", path)

	return &Directory{
		q: q,
		c: r.c,
	}
}

// A schema extension provided by a project
type Extension struct {
	q *querybuilder.Selection
	c graphql.Client
}

// path to the extension's code within the project's filesystem
func (r *Extension) Path(ctx context.Context) (string, error) {
	q := r.q.Select("path")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// schema contributed to the project by this extension
func (r *Extension) Schema(ctx context.Context) (string, error) {
	q := r.q.Select("schema")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// sdk used to generate code for and/or execute this extension
func (r *Extension) Sdk(ctx context.Context) (string, error) {
	q := r.q.Select("sdk")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// A file
type File struct {
	q *querybuilder.Selection
	c graphql.Client
}

// The contents of the file
func (r *File) Contents(ctx context.Context) (string, error) {
	q := r.q.Select("contents")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// The content-addressed identifier of the file
func (r *File) ID(ctx context.Context) (FileID, error) {
	q := r.q.Select("id")

	var response FileID
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

func (r *File) Secret() *Secret {
	q := r.q.Select("secret")

	return &Secret{
		q: q,
		c: r.c,
	}
}

// The size of the file, in bytes
func (r *File) Size(ctx context.Context) (int, error) {
	q := r.q.Select("size")

	var response int
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// A git ref (tag or branch)
type GitRef struct {
	q *querybuilder.Selection
	c graphql.Client
}

// The digest of the current value of this ref
func (r *GitRef) Digest(ctx context.Context) (string, error) {
	q := r.q.Select("digest")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// The filesystem tree at this ref
func (r *GitRef) Tree() *Directory {
	q := r.q.Select("tree")

	return &Directory{
		q: q,
		c: r.c,
	}
}

// A git repository
type GitRepository struct {
	q *querybuilder.Selection
	c graphql.Client
}

// Details on one branch
func (r *GitRepository) Branch(name string) *GitRef {
	q := r.q.Select("branch")
	q = q.Arg("name", name)

	return &GitRef{
		q: q,
		c: r.c,
	}
}

// List of branches on the repository
func (r *GitRepository) Branches(ctx context.Context) ([]string, error) {
	q := r.q.Select("branches")

	var response []string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// Details on one tag
func (r *GitRepository) Tag(name string) *GitRef {
	q := r.q.Select("tag")
	q = q.Arg("name", name)

	return &GitRef{
		q: q,
		c: r.c,
	}
}

// List of tags on the repository
func (r *GitRepository) Tags(ctx context.Context) ([]string, error) {
	q := r.q.Select("tags")

	var response []string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// A set of scripts and/or extensions
type Project struct {
	q *querybuilder.Selection
	c graphql.Client
}

// other projects with schema this project depends on
func (r *Project) Dependencies(ctx context.Context) ([]Project, error) {
	q := r.q.Select("dependencies")

	var response []Project
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// extensions in this project
func (r *Project) Extensions(ctx context.Context) ([]Extension, error) {
	q := r.q.Select("extensions")

	var response []Extension
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// Code files generated by the SDKs in the project
func (r *Project) GeneratedCode() *Directory {
	q := r.q.Select("generatedCode")

	return &Directory{
		q: q,
		c: r.c,
	}
}

// install the project's schema
func (r *Project) Install(ctx context.Context) (bool, error) {
	q := r.q.Select("install")

	var response bool
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// name of the project
func (r *Project) Name(ctx context.Context) (string, error) {
	q := r.q.Select("name")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// schema provided by the project
func (r *Project) Schema(ctx context.Context) (string, error) {
	q := r.q.Select("schema")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// scripts in this project
func (r *Project) Scripts(ctx context.Context) ([]Script, error) {
	q := r.q.Select("scripts")

	var response []Script
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

type Query struct {
	q *querybuilder.Selection
	c graphql.Client
}

// Construct a cache volume from its ID
func (r *Query) Cache(id CacheID) *CacheVolume {
	q := r.q.Select("cache")
	q = q.Arg("id", id)

	return &CacheVolume{
		q: q,
		c: r.c,
	}
}

// Create a new cache volume identified by an arbitrary set of tokens
func (r *Query) CacheFromTokens(tokens []string) *CacheVolume {
	q := r.q.Select("cacheFromTokens")
	q = q.Arg("tokens", tokens)

	return &CacheVolume{
		q: q,
		c: r.c,
	}
}

// ContainerOpts contains options for Query.Container
type ContainerOpts struct {
	ID ContainerID
}

// Load a container from ID.
// Null ID returns an empty container (scratch).
func (r *Query) Container(opts ...ContainerOpts) *Container {
	q := r.q.Select("container")
	// `id` optional argument
	for i := len(opts) - 1; i >= 0; i-- {
		if !querybuilder.IsZeroValue(opts[i].ID) {
			q = q.Arg("id", opts[i].ID)
			break
		}
	}

	return &Container{
		q: q,
		c: r.c,
	}
}

// DirectoryOpts contains options for Query.Directory
type DirectoryOpts struct {
	ID DirectoryID
}

// Load a directory by ID. No argument produces an empty directory.
func (r *Query) Directory(opts ...DirectoryOpts) *Directory {
	q := r.q.Select("directory")
	// `id` optional argument
	for i := len(opts) - 1; i >= 0; i-- {
		if !querybuilder.IsZeroValue(opts[i].ID) {
			q = q.Arg("id", opts[i].ID)
			break
		}
	}

	return &Directory{
		q: q,
		c: r.c,
	}
}

// Load a file by ID
func (r *Query) File(id FileID) *File {
	q := r.q.Select("file")
	q = q.Arg("id", id)

	return &File{
		q: q,
		c: r.c,
	}
}

// Query a git repository
func (r *Query) Git(url string) *GitRepository {
	q := r.q.Select("git")
	q = q.Arg("url", url)

	return &GitRepository{
		q: q,
		c: r.c,
	}
}

// An http remote
func (r *Query) HTTP(url string) *File {
	q := r.q.Select("http")
	q = q.Arg("url", url)

	return &File{
		q: q,
		c: r.c,
	}
}

// Look up a project by name
func (r *Query) Project(name string) *Project {
	q := r.q.Select("project")
	q = q.Arg("name", name)

	return &Project{
		q: q,
		c: r.c,
	}
}

// Load a secret from its ID
func (r *Query) Secret(id SecretID) *Secret {
	q := r.q.Select("secret")
	q = q.Arg("id", id)

	return &Secret{
		q: q,
		c: r.c,
	}
}

// An executable script that uses the project's dependencies and/or extensions
type Script struct {
	q *querybuilder.Selection
	c graphql.Client
}

// path to the script's code within the project's filesystem
func (r *Script) Path(ctx context.Context) (string, error) {
	q := r.q.Select("path")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// sdk used to generate code for and/or execute this script
func (r *Script) Sdk(ctx context.Context) (string, error) {
	q := r.q.Select("sdk")

	var response string
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}

// A reference to a secret value, which can be handled more safely than the value itself
type Secret struct {
	q *querybuilder.Selection
	c graphql.Client
}

// The identifier for this secret
func (r *Secret) ID(ctx context.Context) (SecretID, error) {
	q := r.q.Select("id")

	var response SecretID
	q = q.Bind(&response)
	return response, q.Execute(ctx, r.c)
}
