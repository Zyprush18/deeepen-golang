<script lang="ts">
	const prop = $props();
	let data = $state({
		from: "",
		name: "",
		message: "",
		to: "",
		type: ""
	});

	
	if (prop.msg === '' || prop.msg != null) {
		data = JSON.parse(prop.msg);
	}	
</script>

<!-- {@debug name} -->
{#if data.from === prop.slug || data.to === prop.slug}
	<div
		class="p-3"
		class:text-right={data.name === prop.profile.data.username}
		class:text-left={data.name !== prop.profile.data.username}
	>
			<span class="text-sm text-gray-500">{data.from === prop.slug ? prop.friend.name: "Anda"}</span>
		<p class="mb-4">{data.message}</p>
	</div>
{/if}
