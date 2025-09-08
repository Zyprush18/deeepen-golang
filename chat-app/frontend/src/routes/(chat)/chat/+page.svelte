<script lang="ts">
	import AppLayout from '$lib/components/layout/App-Layout.svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>ZHAT</title>
</svelte:head>

{#snippet test()}
	<div class="flex h-full items-center justify-center">
		<h1>Hello Zhat</h1>
	</div>
{/snippet}

<AppLayout slug={null} profile={data.Profile} children={test} />
