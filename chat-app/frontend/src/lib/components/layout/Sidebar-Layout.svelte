<script lang="ts">
    import ChevronUp from "@lucide/svelte/icons/chevron-up";
    import * as Sidebar from "$lib/components/ui/sidebar/index.js";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import profile3 from "$lib/assets/avatars/3.png"
    import profile6 from "$lib/assets/avatars/6.png"
    import profile10 from "$lib/assets/avatars/10.png"
    import profile23 from "$lib/assets/avatars/23.png"
    import profile26 from "$lib/assets/avatars/26.png"
    import profile25 from "$lib/assets/avatars/25.png"
    

  // Menu items.
    const items = [
        {
            title: "Angga",  
            url: "#",
            src : profile3
        },
        {
            title: "Farel",
            url: "#",
            src: profile6
        },
        {
            title: "Muaiman",
            url: "#",
            src: profile10
        },
        {
            title: "Eko",
            url: "#",
            src: profile26
        },
        {
            title: "Budi",
            url: "#",
            src: profile25
        },
    ];
</script>

<Sidebar.Root>
    <Sidebar.Header>
        <h1 class="text-4xl p-3 font-bold text-center text-green-600">ZHAT</h1>
    </Sidebar.Header>

    <Sidebar.Content>
        <Sidebar.Group>
        <Sidebar.GroupLabel>Message</Sidebar.GroupLabel>
        <Sidebar.GroupContent>
            <Sidebar.Menu>
            {#each items as item (item.title)}
                <Sidebar.MenuItem>
                <Sidebar.MenuButton class="py-8">
                    {#snippet child({ props })}
                    <a href={item.url} {...props}>
                        <img src={item.src} alt="" class="w-[47px] h-[47px]">
                        <span class="text-lg">{item.title}</span>
                    </a>
                    {/snippet}
                </Sidebar.MenuButton>
                </Sidebar.MenuItem>
            {/each}
            </Sidebar.Menu>
        </Sidebar.GroupContent>
        </Sidebar.Group>
    </Sidebar.Content>

    <Sidebar.Footer>
        <Sidebar.Menu>
        <Sidebar.MenuItem>
            <DropdownMenu.Root>
            <DropdownMenu.Trigger>
                {#snippet child({ props })}
                <Sidebar.MenuButton
                    {...props}
                    class="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground py-8 text-md font-bold"
                >
                    <img src={profile23} alt="" class="w-[47px] h-[47px] mr-2">
                    Zul
                    <ChevronUp class="ml-auto" />
                </Sidebar.MenuButton>
                {/snippet}
            </DropdownMenu.Trigger>
            <DropdownMenu.Content
                side="top"
                class="w-(--bits-dropdown-menu-anchor-width)"
            >
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
