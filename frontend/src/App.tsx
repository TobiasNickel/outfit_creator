

import { useState } from 'react'
import * as outfitSdk from './sdk/outfit'
import './App.css'
import { Item } from './components/Item.component';
import { Gender } from './sdk/outfit';

export function App() {
    const [loading, setLoading] = useState(false)
    const [outfit, setOutfit] = useState(undefined as outfitSdk.TItemResponse | undefined);
    const [error, setError] = useState(undefined as any);
    const [gender, setGender] = useState('FEMALE' as Gender)
    const [preferImageType, setPreferImageType] = useState(outfitSdk.ImageType.OUTFIT_IMAGE)
    const [preferLanguage, setPreferLanguage] = useState(outfitSdk.Language.DE)

    async function loadRandomOutfit(loadGender: Gender) {
        setLoading(true);
        setOutfit(undefined)
        try {
            const loadedOutfit = await outfitSdk.random(loadGender, preferLanguage);
            setOutfit(loadedOutfit)
            console.log(loadedOutfit)
        } catch (err) {
            setError(err);
        }
        setLoading(false)
    }
    return (
        <>
            <h1>Outfit Vorschlag</h1>
            <div className='controlls'>
                <select value={gender} onChange={(e)=>setGender(e.target.value as Gender)}>
                    <option value='FEMALE'>👩‍🦰 weiblich</option>
                    <option value='MALE'>👨‍🦰männlich</option>
                </select>
                <button onClick={() => { loadRandomOutfit(gender) }} disabled={loading}>presentiere Outfit</button>
                {loading ? <div className='loading'>loading...</div> : <></>}
                {error ? <>{error}</> : <></>}
            </div>
            {outfit?<>
                <Item item={outfit.upper} imageType={preferImageType} language={preferLanguage}></Item>
                <Item item={outfit.under} imageType={preferImageType} language={preferLanguage}></Item>
                <Item item={outfit.accessory} imageType={preferImageType} language={preferLanguage}></Item>
            </>:<></>}
            <div>
                <h2>Einstellungen</h2>
                <div>
                    <span className="label">Bildtype</span>
                    <select value={preferImageType} onChange={(e)=>setPreferImageType(e.target.value as outfitSdk.ImageType)}>
                        <option value={outfitSdk.ImageType.CUTOUT}>Product</option>
                        <option value={outfitSdk.ImageType.OUTFIT_IMAGE}>Beispiel</option>
                    </select>
                </div>
                <div>
                    <span className="label">Sprache</span>
                    <select value={preferLanguage} onChange={(e)=>setPreferLanguage(e.target.value as outfitSdk.Language)}>
                        <option value={outfitSdk.Language.BS}>BS</option>
                        <option value={outfitSdk.Language.CS}>CS</option>
                        <option value={outfitSdk.Language.DE}>DE</option>
                        <option value={outfitSdk.Language.EN}>EN</option>
                        <option value={outfitSdk.Language.HR}>HR</option>
                        <option value={outfitSdk.Language.SR}>SR</option>
                        <option value={outfitSdk.Language.RU}>RU</option>
                        <option value={outfitSdk.Language.PL}>PL</option>
                        <option value={outfitSdk.Language.BS}>BS</option>
                        <option value={outfitSdk.Language.BS}>BS</option>
                        <option value={outfitSdk.Language.BE}>BE</option>
                        <option value={outfitSdk.Language.SL}>SL</option>
                    </select>
                </div>
            </div>
        </>
    )
}
