<script lang="ts">
	import { onMount } from 'svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import AppLayout from '$lib/components/layout/App-Layout.svelte';
	import Messages from '$lib/components/Messages.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { SendHorizontal } from '@lucide/svelte';
	import wss from '$lib/utils/storews.js';

	let message = $state('');
	let messages = $state<string[]>([]);
	let { data } = $props();

	onMount(() => {
		wss.connect(data?.Auth);
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
	{#each messages as msg, i}
		<div class="p-7">
			{#if msg !== ''}
				<Messages {msg} direction={i % 2 == 0 ? 'right' : 'left'} />
			{/if}
		</div>
	{/each}

	<div class="flex min-h-full justify-center">
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

<AppLayout children={test} />
