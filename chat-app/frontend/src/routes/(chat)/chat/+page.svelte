<script lang="ts">
	import AppLayout from '$lib/components/layout/App-Layout.svelte';

	let {data} = $props();
	console.log(data?.Auth);
	
</script>

<svelte:head>
	<title>ZHAT</title>
</svelte:head>

{#snippet test()}
	<div class="h-full p-5">
		<h1>Kita</h1>
	</div>
{/snippet}

<AppLayout children={test} />
