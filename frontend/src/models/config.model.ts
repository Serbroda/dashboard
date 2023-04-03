export type Config = {
    title: string;
    sections: Section[];
}

export type Section = {
    name: string;
    items: Item[];
}

export type Item = {
    name: string;
    url: string
}
