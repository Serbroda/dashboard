<script lang="ts">
    import { onDestroy } from 'svelte';
    import type {Item} from "../models/config.model";

    export let item: Item = {name: "", url: ""};
    let online: boolean = false;
    let interval: number | undefined;

    const check = (url: string): Promise<boolean> => {
        return new Promise((res, rej) => {
            console.debug('Checking availability for ', url);

            fetch(url, {mode: 'no-cors'}).then((r) => {
                console.debug('result for url ' + url, r);
                return res(true);
            }).catch(err => {
                console.debug('result for url ' + url, err);
                return res(false);

            });
        });
    }

    if (item.checkAvailability === true) {
        check(item.url).then(res => online = res);

        if (item.checkAvailabilityIntervalSeconds) {
            console.debug('Interval', {
                url: item.url,
                seconds: item.checkAvailabilityIntervalSeconds
            });
            interval = setInterval(() => {
                check(item.url).then(res => online = res);
            }, 1000 * item.checkAvailabilityIntervalSeconds)
        }
    }

    onDestroy(() => clearInterval(interval));
</script>

<a target="_blank" style="text-decoration: none" class="flex bg-surface-700 p-2 rounded-md"
   href={item.url}>{item.name}
    {#if (item.checkAvailability === true)}
        <div class="flex-1"></div>
        <div class={`w-2 self-start aspect-square rounded-full animate-pulse ${online ? "bg-green-500" : "bg-red-500"}`}></div>
    {/if}
</a>