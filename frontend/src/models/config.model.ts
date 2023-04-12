export type Config = {
    title: string;
    description?: string;
    sections: Section[];
}

export type Section = {
    name: string;
    items: Item[];
}

export type Item = {
    name: string;
    description?: string;
    url: string;
    icon?: string;
    checkAvailability?: boolean;
    checkAvailabilityIntervalSeconds?: number;
}
