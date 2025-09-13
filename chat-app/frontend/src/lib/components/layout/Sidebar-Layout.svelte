<script lang="ts">
	import ChevronUp from '@lucide/svelte/icons/chevron-up';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import profile10 from '$lib/assets/avatars/10.png';
	import profile23 from '$lib/assets/avatars/23.png';
	import { Plus } from '@lucide/svelte';
	import * as Dialog from '$lib/components/ui/dialog/index.js';

	const { profile, friends, create } = $props();
	const friendPending = friends.data.filter((f: { status: string }) => f.status == 'pending');
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
			<Sidebar.GroupLabel class="justify-between">
				<span>Message</span>
				<Dialog.Root>
					<Dialog.Trigger ><Plus /></Dialog.Trigger>
					<Dialog.Content class="sm:max-w-[425px]">
						<Dialog.Header>
							<Dialog.Title>Add</Dialog.Title>
							<Dialog.Description>
								Added Your Friends
						</Dialog.Description>
						</Dialog.Header>

						{@render create?.()}
						
					</Dialog.Content>
				</Dialog.Root>
			</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each friends as item}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton class="py-8">
								{#snippet child({ props })}
									<a href={`/chat/${item.uuid}`} {...props}>
										<img src={profile10} alt="" class="h-[47px] w-[47px]" />
										<span class="text-lg"
											>{item.role === 'received' ? item.name_from : item.name_to}</span
										>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>

		{#if friendPending != null}
			<Sidebar.Group>
				<Sidebar.GroupLabel>Message Pending</Sidebar.GroupLabel>
				<Sidebar.GroupContent>
					<Sidebar.Menu>
						{#each friendPending as item}
							<Sidebar.MenuItem>
								<Sidebar.MenuButton class="py-8">
									{#snippet child({ props })}
										<a href={`/chat/${item.uuid}`} {...props}>
											<img src={profile10} alt="" class="h-[47px] w-[47px]" />
											<span class="text-lg"
												>{item.role === 'received'
													? item.name_from !== ''
														? item.name_from
														: '~' + item.uuid
													: item.name_to !== ''
														? item.name_to
														: '~' + item.uuid}</span
											>
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
								{profile.data.username}
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
