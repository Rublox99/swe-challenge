import type { Hit, ZincResponse } from '@/interfaces/Zinc'
import axios from 'axios'

class ZincService {
    private static localhostURL = "http://localhost:8080/emails"
    private static ec2URL = "http://ec2-3-144-1-53.us-east-2.compute.amazonaws.com:8080/emails"
    private static railwayURL = "https://swe-backend-production.up.railway.app/emails"

    static async GetFilteredEmails(from: number, size: number, startDate: string, endDate: string, text: string): Promise<ZincResponse> {
        try {
            const response = await axios.get<ZincResponse>(`${this.railwayURL}/filters`, {
                params: {
                    from, size,
                    start_date: startDate,
                    end_date: endDate,
                    text
                }
            })

            return response.data
        } catch (error) {
            throw new Error(`Failed to fetch the filtered emails: ${error}`)
        }
    }

    static async GetAllEmails(from: number, size: number): Promise<ZincResponse> {
        try {
            const response = await axios.get<ZincResponse>(`${this.railwayURL}/all`, {
                params: {
                    from, size
                }
            })

            return response.data
        } catch (error) {
            throw new Error(`Failed to fetch all the emails: ${error}`)
        }
    }

    static async GetEmailById(id: string): Promise<Hit> {
        try {
            const response = await axios.get<Hit>(`${this.railwayURL}/index`, {
                params: {
                    id
                }
            })

            return response.data
        } catch (error) {
            throw new Error(`Failed to fetch the email by ID: ${error}`)
        }
    }
}

export {
    ZincService
}