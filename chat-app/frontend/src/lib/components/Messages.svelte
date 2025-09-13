<script lang="ts">
	const prop = $props();
	let data = $state({
		from: '',
		name: '',
		message: '',
		to: '',
		type: ''
	});

	if (prop.msg === '' || prop.msg != null) {
		data = JSON.parse(prop.msg);
	}

	console.log(prop.friend.data);
	
</script>

<!-- {@debug data} -->
{#if data.from === prop.slug || data.to === prop.slug}
	<div
		class="m-4 flex"
		class:justify-end={data.name === prop.profile.data.username}
		class:justify-start={data.name !== prop.profile.data.username}
	>
		<div class="max-w-[75%] rounded-lg bg-white p-3 shadow">
			{#each prop.friend.data as item}
				<span class="mb-1 block text-xs text-gray-600">
					{data.from === prop.slug
						? item.role === 'received'
							? item.name_from
							:'~'+item.name_to || item.uuid
						: 'Anda'}
				</span>
			{/each}
			<p class="text-sm text-gray-800">{data.message}</p>
		</div>
	</div>
{/if}
