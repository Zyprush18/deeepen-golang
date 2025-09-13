<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageProps } from './$types';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import AppLayout from '$lib/components/layout/App-Layout.svelte';
	import Messages from '$lib/components/Messages.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { SendHorizontal } from '@lucide/svelte';
	import wss from '$lib/utils/storews';
	import CreateFriends from '$lib/components/Create.svelte';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';


	let message = $state('');
	let messages = $state<string[]>([]);
	let { data }: PageProps = $props();

	$effect(() => {
		wss.connect(data?.Auth, data?.Slug);
	});

	onMount(() => {
		wss.subscribe((msg) => {
			messages = [...messages, msg];
		});
	});
	function sendMessage() {
		if (message.length > 0) {
			wss.send(message);
			message = '';
		}
	}

</script>

<svelte:head>
	<title>ZHAT</title>
</svelte:head>

{#snippet test()}
	{#each messages as msg}
		{#if msg !== ''}
			<Messages {msg} profile={data.Profile} slug={data.Slug} friend={data.Friend}  />
		{/if}
	{/each}

		<div class="sticky flex justify-center mt-14">
			<div class="fixed bottom-0 mb-4 flex w-full justify-center space-x-3">
				<Textarea
					name="message"
					bind:value={message}
					placeholder="enter your message....."
					class="max-h-40 min-h-[40px] w-[60%] resize-none overflow-y-auto bg-gray-100 shadow-md"
				/>
				<Button class="h-10 w-10 rounded-full" onclick={sendMessage}><SendHorizontal /></Button>
			</div>
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
<AppLayout slug={data.Slug} profile={data.Profile} children={test} friends={data.Friend} createfriend={create} />
