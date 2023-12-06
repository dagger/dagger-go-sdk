defmodule Dagger do
  @moduledoc """
  The [Dagger](https://dagger.io/) SDK for Elixir.

  See `getting_start.livemd` for starter point.
  """

  defstruct [:client, :query]

  @doc """
  Connecting to Dagger.

  ## Options

  #{NimbleOptions.docs(Dagger.Internal.Client.connect_schema())}
  """
  def connect(opts \\ []) do
    with {:ok, graphql_client} <- Dagger.Internal.Client.connect(opts) do
      client = %Dagger.Client{
        client: graphql_client,
        selection: Dagger.QueryBuilder.Selection.query()
      }

      case Dagger.Client.check_version_compatibility(client, Dagger.EngineConn.engine_version()) do
        {:error, reason} -> IO.warn("failed to check version compatibility: #{inspect(reason)}")
        _ -> nil
      end

      {:ok, client}
    end
  end

  @doc """
  Similar to `connect/1` but raise exception when found an error.
  """
  def connect!(opts \\ []) do
    case connect(opts) do
      {:ok, query} -> query
      error -> raise "Cannot connect to Dagger engine, cause: #{inspect(error)}"
    end
  end

  @doc """
  Disconnecting Dagger.
  """
  def close(%Dagger.Client{client: client}) do
    Dagger.Internal.Client.close(client)
  end
end