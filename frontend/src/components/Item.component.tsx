import { ImageType, Language, TItem } from "../sdk/outfit"
import './Item.css'

export function Item({item, imageType, language}: {item: TItem, imageType: ImageType, language: Language}) {
    if(!item){
        return (
            <>no item</>
        )
    }
    const images = findAllItemImages(item)
    const displayImage = images.find(i=>i.type==imageType)||images[0]
    const actualImageType = displayImage.type==ImageType.CUTOUT?'freigestellt':'Beispiel'
    return (
        <div className="item">
            <div className="name">{item.descriptions.find(d=>d.language===language)?.description}</div>
            <div className="brand">{item.brand}</div>
            <div className="maintenance_group">{item.maintenance_group}</div>
            <div className="price">{getPrice(item)}</div>
            <div>{item.id}</div>
            <div className="buy_now_link"><a href={`https://www.newyorker.de/products/#/detail/${item.id}/001`} target="_blank">jetzt einpacken</a></div>
            <img src={`https://api.newyorker.de/csp/images/image/public/${displayImage.key}?res=high&frame=1_1`} alt={actualImageType}></img>
        </div>
    )
}

function findAllItemImages(item: TItem){
    return item.variants.map(v=>v.images).flat();
}

function getPrice(item: TItem): string{
    const min = Math.min(...item.variants.map(v=>v.current_price));
    const max = Math.max(...item.variants.map(v=>v.current_price));
    var currency = item.variants[0].currency
    if(min===max){
        return min+' '+currency
    }
    return min+' - '+max+' '+currency
}