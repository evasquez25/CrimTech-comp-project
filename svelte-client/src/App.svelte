<script>
    import { onMount } from "svelte";
    import { IncrementRequest } from "./proto/counter/counter_pb.js";
    import { CounterServiceClient } from "./proto/counter/counter_grpc_web_pb.js";
    import { grpc } from "@improbable-eng/grpc-web";

    let counter = 0;
    let error = null;
    let loading = false;

    const client = new CounterServiceClient("http://localhost:8080", {
        transport: grpc.WebsocketTransport()
    });

    async function incrementCounter() {
        const request = new IncrementRequest();
        error = null;
        loading = true;

        try {
            const metadata = {};
            client.incrementCounter(request, metadata, (err, response) => {
                loading = false;
                if (err) {
                    console.error("Error:", err);
                    error = err.message || "Failed to increment counter";
                    return;
                }
                counter = response.getValue();
            });
        } catch (e) {
            loading = false;
            error = e.message || "Failed to increment counter";
            console.error("Error:", e);
        }
    }

    onMount(() => {
        console.log("Component mounted");
    });
</script>

<main>
    <h1>Counter</h1>
    <p>Current Value: {counter}</p>
    {#if error}
        <p class="error">{error}</p>
    {/if}
    <button on:click={incrementCounter} disabled={loading}>
        {loading ? 'Incrementing...' : 'Increment Counter'}
    </button>
</main>

<style>
    main {
        text-align: center;
        padding: 20px;
        font-family: sans-serif;
        max-width: 600px;
        margin: 0 auto;
    }
    button {
        padding: 10px 20px;
        font-size: 18px;
        cursor: pointer;
        background-color: #4CAF50;
        color: white;
        border: none;
        border-radius: 4px;
        transition: background-color 0.3s;
    }
    button:hover:not(:disabled) {
        background-color: #45a049;
    }
    button:disabled {
        background-color: #cccccc;
        cursor: not-allowed;
    }
    .error {
        color: #ff3e3e;
        margin: 10px 0;
    }
</style>