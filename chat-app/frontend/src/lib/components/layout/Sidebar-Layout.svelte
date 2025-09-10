<script lang="ts">
	import ChevronUp from '@lucide/svelte/icons/chevron-up';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import profile10 from '$lib/assets/avatars/10.png';
	import profile23 from '$lib/assets/avatars/23.png';

	const { data } = $props();
	const acceptData = data.data.friend.filter((f: { status: string; }) => f.status === 'accept')
	const pendingData = data.data.friend.filter((f: { status: string; }) => f.status === 'pending')
	
</script>

<Sidebar.Root>
	<Sidebar.Header>
		<h1
			class="m-3 scroll-m-20 text-center text-3xl font-extrabold tracking-tight text-green-600 lg:text-5xl"
		>
			ZHAT
		</h1>
	</Sidebar.Header>

	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupLabel>Message</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each acceptData as item}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton class="py-8">
								{#snippet child({ props })}
									<a href={`/chat/${item.uuid}`} {...props}>
										<img src={profile10} alt="" class="h-[47px] w-[47px]" />
										<span class="text-lg">{item.name}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>

		{#if pendingData.find((f: {status: string})=> f.status === "pending")}
			<Sidebar.Group>
			<Sidebar.GroupLabel>Message Pending</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each pendingData as item}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton class="py-8">
								{#snippet child({ props })}
									<a href={`/chat/${item.uuid}`} {...props}>
										<img src={profile10} alt="" class="h-[47px] w-[47px]" />
										<span class="text-lg">{item.name}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
		{/if}
	</Sidebar.Content>

	<Sidebar.Footer>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<DropdownMenu.Root>
					<DropdownMenu.Trigger>
						{#snippet child({ props })}
							<Sidebar.MenuButton
								{...props}
								class="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground text-md py-8 font-bold"
							>
								<img src={profile23} alt="" class="mr-2 h-[47px] w-[47px]" />
								{data.data.username}
								<ChevronUp class="ml-auto" />
							</Sidebar.MenuButton>
						{/snippet}
					</DropdownMenu.Trigger>
					<DropdownMenu.Content side="top" class="w-(--bits-dropdown-menu-anchor-width)">
						<a href="/login">
							<DropdownMenu.Item>
								<span>Sign out</span>
							</DropdownMenu.Item>
						</a>
					</DropdownMenu.Content>
				</DropdownMenu.Root>
			</Sidebar.MenuItem>
		</Sidebar.Menu>
	</Sidebar.Footer>
</Sidebar.Root>
