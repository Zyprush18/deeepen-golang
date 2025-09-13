<script lang="ts">
    import { buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';


    const {content } =$props()
</script>


<Dialog.Root>
	<Dialog.Trigger class={buttonVariants({ variant: 'outline' })}>Add Friends</Dialog.Trigger>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>Add Friends</Dialog.Title>
			<Dialog.Description>
				Add your friend for chatt
			</Dialog.Description>
		</Dialog.Header>
		{@render content?.()}
	</Dialog.Content>
</Dialog.Root>
