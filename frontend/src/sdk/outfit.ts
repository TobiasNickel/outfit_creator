
export async function random(gender: Gender, language: Language): Promise<TItemResponse> {
    return fetch('/api/outfit/random?gender=' + encodeURIComponent(gender) + '&country=' + encodeURIComponent(language)).then(r => r.json());
}

export enum Gender {
    male = "MALE",
    female = "FEMALE",
}

export interface TItemResponse {
    upper: TItem
    under: TItem
    accessory: TItem
}

export interface TItem {
    global_item_id: string
    id: string
    country: string
    maintenance_group: string
    web_category_id: string
    web_category: string
    brand: string
    sales_unit: number
    customer_group: string
    variants: Variant[]
    descriptions: Description[]
}

interface Variant {
    id: string
    product_id: string
    publish_date: string
    new_in: boolean
    coming_soon: boolean
    sale: boolean
    highlight: boolean
    color_name: string
    pantone_color: string
    pantone_color_name: string
    red: number
    green: number
    blue: number
    color_group: string
    basic_color: string
    currency: string
    original_price: number
    current_price: number
    red_price_change: boolean
    sizes: Size[]
    images: Image[]
}

interface Size {
    size_value: string
    size_name: string
    bar_code: string
}

interface Image {
    key: string
    type: ImageType
    angle: string
    has_thumbnail: boolean
    position: number
}

export enum ImageType {
    OUTFIT_IMAGE = "OUTFIT_IMAGE",
    CUTOUT = "CUTOUT",
}

interface Description {
    language: string
    description: string
}

export enum Language {
    BS = "BS",
    CS = "CS",
    DE = "DE",
    EN = "EN",
    HR = "HR",
    SR = "SR",
    RU = "RU",
    PL = "PL",
    BE = "BE",
    SL = "SL",
}