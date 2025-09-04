<script>
	import AppLayout from '$lib/components/layout/App-Layout.svelte';
</script>

{#snippet test()}
	<!-- <div class="flex h-full items-center justify-center">ddlklkk</div> -->
    <h1 class="text-2xl">name</h1>
{/snippet}

<AppLayout children={test} />
