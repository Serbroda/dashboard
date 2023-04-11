<script lang="ts">
    import {onDestroy} from 'svelte';
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

<a target="_blank" class="app-item" href={item.url}>
    {#if item.icon}
        <div class="self-center">
            <img class="h-8 w-8" src={item.icon}/>
        </div>
    {/if}
    <div class="flex-1 flex flex-col pl-2">
        <span>{item.name}</span>
        <small>Hi</small>
    </div>
    <div>
        {#if (item.checkAvailability === true)}
            <div class="flex-1"></div>
            <div class={`app-item-availability ${online ? "app-item-availability-online" : "app-item-availability-offline"}`}></div>
        {/if}
    </div>
</a>
