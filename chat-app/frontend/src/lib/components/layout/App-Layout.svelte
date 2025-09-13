<script lang="ts">
    import * as Sidebar from "$lib/components/ui/sidebar/index.js";
    import AppSidebar from "$lib/components/layout/Sidebar-Layout.svelte";
    import Navbar from "$lib/components/layout/Navbar-Layout.svelte";
	import Bgchat from '$lib/assets/wabg.png';

    let { children,profile,slug,friends,createfriend } = $props();
    
</script>

<Sidebar.Provider class="min-h-screen flex">
    <AppSidebar profile={profile} friends={friends} create={createfriend} />
    <main class="flex-1 block"  style={`background-image: url(${Bgchat})`}>
        <Navbar profil={friends} slug={slug} />
        {@render children?.()}
    </main>
</Sidebar.Provider>
