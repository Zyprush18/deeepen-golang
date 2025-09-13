<script lang="ts">
	import AppLayout from '$lib/components/layout/App-Layout.svelte';
	import wsss from '$lib/utils/storews.js';
	import CreateFriends from '$lib/components/Create.svelte';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
    import { Button } from '$lib/components/ui/button/index.js';


	let { data } = $props();

	$effect(() => {
		wsss.connect(data?.Auth, null);
	});
</script>

<svelte:head>
	<title>ZHAT</title>
</svelte:head>

{#snippet welcome()}
	<div class="flex h-full items-center justify-center ">
		<h1 class="text-7xl font-bold text-gray-400">Welcome To Zhat</h1>
	</div>
{/snippet}



{#snippet create()}
	{#snippet field()}
		<div class="grid gap-4 py-4">
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="name_from" class="text-right">Name</Label>
				<Input id="name_from" name="name_from" placeholder="Pedro Duarte" class="col-span-3" />
			</div>
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="email" class="text-right">Email</Label>
				<Input id="email" name="email" placeholder="email@gmail.com" class="col-span-3" />
			</div>
		</div>
		<Dialog.Footer>
			<Button type="submit">Add</Button>
		</Dialog.Footer>
	{/snippet}

	<CreateFriends content={field} />
{/snippet}

<AppLayout
	slug={null}
	profile={data.Profile}
	children={welcome}
	friends={data.Friend}
	createfriend={create}
/>
