<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import ZodIssues from '$lib/components/ZodIssues.svelte';
	let { form } = $props();

	type formsResp = {
		data: any;
	};

	const response: formsResp = {
		data: form
	};
</script>

<svelte:head>
	<title>Register</title>
</svelte:head>

<!-- <h1 class="text-5xl font-bold text-green-600 text-center">ZHAT</h1> -->
<div class="flex h-screen items-center justify-center">
	{#if response.data != null}
		<ZodIssues issues={response.data} />
	{/if}
	<Card.Root class="w-full max-w-sm shadow-2xl">
		<Card.Header class="text-center">
			<span class="text-lg font-bold text-green-600">ZHAT</span>
			<Card.Title class="text-2xl">Register</Card.Title>
			<Card.Description>Enter your email below to create a new account</Card.Description>
		</Card.Header>
		<form method="POST">
			<Card.Content>
				<div class="mb-3 flex flex-col gap-6">
					<div class="grid gap-2">
						<Label for="username">Username</Label>
						<Input id="username" type="text" name="username" placeholder="Joko" />
					</div>
					<div class="grid gap-2">
						<Label for="email">Email</Label>
						<Input id="email" type="email" name="email" placeholder="m@example.com" />
					</div>
					<div class="grid gap-2">
						<div class="flex items-center">
							<Label for="password">Password</Label>
						</div>
						<Input
							id="password"
							type="password"
							name="password"
							placeholder="enter your password...."
						/>
					</div>
					<div class="grid gap-2">
						<div class="flex items-center">
							<Label for="repassword">Confirm Password</Label>
						</div>
						<Input
							id="repassword"
							type="password"
							name="repassword"
							placeholder="enter your password again...."
						/>
					</div>
				</div>
				<span class="mb-5 ml-auto inline-block text-sm text-gray-500 underline-offset-4">
					Sudah Punya Akun? <a href="/login" class="text-blue-500 hover:underline">Login</a>
				</span>
			</Card.Content>
			<Card.Footer class="flex-col gap-2">
				<Button type="submit" class="w-full  bg-green-500 hover:bg-green-600">Register</Button>
			</Card.Footer>
		</form>
	</Card.Root>
</div>
