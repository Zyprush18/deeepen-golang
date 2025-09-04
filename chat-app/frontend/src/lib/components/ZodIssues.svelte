<script lang="ts">
	import * as Alert from '$lib/components/ui/alert/index.js';
	import AlertCircleIcon from '@lucide/svelte/icons/alert-circle';
	import type { ZodIssue } from 'zod';


    let {issues}: {issues:ZodIssue[]} = $props()

    console.log(issues);
    
    
</script>

<div class="absolute grid w-full max-w-xl items-end top-3 right-4 gap-4 z-30">
	<Alert.Root variant="destructive">
		<AlertCircleIcon />
		<Alert.Title>Failed Register</Alert.Title>
		<Alert.Description>
			<p>Your input doesn't match the required format. Verify each field and resubmit.</p>
			<ul class="list-inside list-disc text-sm">
                    {#each issues as items }
                        <li>{items.message}</li>
                    {/each}

			</ul>
		</Alert.Description>
	</Alert.Root>
</div>
